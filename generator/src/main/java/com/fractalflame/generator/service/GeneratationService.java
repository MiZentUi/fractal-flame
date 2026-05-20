package com.fractalflame.generator.service;

import java.util.UUID;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.TimeUnit;

import org.springframework.stereotype.Service;

import com.fractalflame.generator.exception.GeneratorException;
import com.fractalflame.generator.model.GenerationTask;
import com.fractalflame.generator.properties.GeneratorProperties;
import com.fractalflame.generator.repository.FractalsRepository;
import com.fractalflame.generator.utils.Generator;

import jakarta.annotation.PostConstruct;
import jakarta.annotation.PreDestroy;
import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class GeneratationService {
    private final GeneratorProperties properties;
    private final Generator generator;
    private final StorageService storageService;
    private final FractalsRepository fractalsRepository;

    private final BlockingQueue<GenerationTask> taskQueue = new LinkedBlockingQueue<>();

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
    }
}
