package com.fractalflame.generator.service;

import java.util.Map;
import java.util.UUID;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.TimeUnit;

import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Service;

import com.fractalflame.generator.exception.GeneratorException;
import com.fractalflame.generator.exception.TaskNotFoundException;
import com.fractalflame.generator.properties.GeneratorProperties;
import com.fractalflame.generator.proto.TaskState;
import com.fractalflame.generator.repository.FractalsRepository;
import com.fractalflame.generator.utils.GenerationTask;
import com.fractalflame.generator.utils.Generator;

import io.grpc.stub.StreamObserver;
import jakarta.annotation.PostConstruct;
import jakarta.annotation.PreDestroy;
import lombok.RequiredArgsConstructor;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@Service
@RequiredArgsConstructor
public class GeneratationService {
    private final GeneratorProperties properties;
    private final Generator generator;
    private final StorageService storageService;
    private final FractalsRepository fractalsRepository;

    private final BlockingQueue<GenerationTask> taskQueue = new LinkedBlockingQueue<>();
    private final Map<Long, GenerationTask> pendingTasks = new ConcurrentHashMap<>();

    private ExecutorService workers;

    @PostConstruct
    void init() {
        workers = Executors.newFixedThreadPool(properties.getCount());

        for (int i = 0; i < properties.getCount(); i++) {
            workers.submit(() -> {
                while (!Thread.currentThread().isInterrupted()) {
                    try {
                        var task = taskQueue.take();
                        generator.process(task);
                        pendingTasks.remove(task.getFractal().getId());
                        sendUpdate(task);

                        var fractal = task.getFractal();
                        fractal.setImage(storageService.saveImage(UUID.randomUUID().toString(),
                                task.getImage()));

                        synchronized (fractalsRepository) {
                            fractalsRepository.save(fractal);
                        }
                    } catch (InterruptedException e) {
                        Thread.currentThread().interrupt();
                        throw new GeneratorException(e);
                    }
                }
            });
        }
    }

    @Scheduled(fixedDelayString = "${app.tasks-update-delay}")
    void tasksUpdate() {
        for (var task : pendingTasks.values()) {
            sendUpdate(task);
        }
    }

    @PreDestroy
    void free() {
        workers.shutdown();
        try {
            if (!workers.awaitTermination(5, TimeUnit.SECONDS)) {
                workers.shutdownNow();
            }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            throw new GeneratorException(e);
        }
    }

    public void addTask(GenerationTask task) {
        taskQueue.add(task);
        pendingTasks.put(task.getFractal().getId(), task);
    }

    public void subscribeToTask(Long id, StreamObserver<TaskState> responseObserver) {
        if (pendingTasks.containsKey(id)) {
            pendingTasks.get(id).setResponseObserver(responseObserver);
        } else {
            throw new TaskNotFoundException(String.format("Task with id=%s not pending!", id));
        }
    }

    private void sendUpdate(GenerationTask task) {
        var responseObserver = task.getResponseObserver();
        if (responseObserver != null) {
            log.atInfo().addKeyValue("fractal_id", task.getFractal().getId()).log("send task update");
            responseObserver.onNext(task.toTaskState());
        }
    }
}
