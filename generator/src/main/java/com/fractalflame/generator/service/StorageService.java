package com.fractalflame.generator.service;

import org.springframework.stereotype.Service;

import com.fractalflame.generator.exception.StorageException;
import com.fractalflame.generator.model.Image;
import com.fractalflame.generator.utils.ImageWriter;

import jakarta.annotation.PostConstruct;
import lombok.RequiredArgsConstructor;
import software.amazon.awssdk.core.sync.RequestBody;
import software.amazon.awssdk.services.s3.S3Client;
import software.amazon.awssdk.services.s3.model.GetObjectRequest;
import software.amazon.awssdk.services.s3.model.NoSuchBucketException;
import software.amazon.awssdk.services.s3.model.PutObjectRequest;

@Service
@RequiredArgsConstructor
public class StorageService {
    private final S3Client s3Client;

    private static final String IMAGES_BUCKET = "images";

    @PostConstruct
    void init() {
        try {
            s3Client.headBucket(request -> request.bucket(IMAGES_BUCKET));
        } catch (NoSuchBucketException exception) {
            s3Client.createBucket(request -> request.bucket(IMAGES_BUCKET));
        }
    }

    public synchronized String saveImage(String filename, Image image) {
        var request = PutObjectRequest.builder()
                .bucket(IMAGES_BUCKET)
                .key(filename)
                .contentType("image/png")
                .build();

        s3Client.putObject(request,
                RequestBody.fromBytes(ImageWriter.toByteArray(image)));

        return filename;
    }

    public byte[] getImageBytes(String filename) {
        var request = GetObjectRequest.builder()
                .bucket(IMAGES_BUCKET)
                .key(filename)
                .build();

        try (var stream = s3Client.getObject(request)) {
            return stream.readAllBytes();
        } catch (Exception e) {
            throw new StorageException("IMAGE_NOT_FOUND", String.format("Image with name %s not found!", filename));
        }
    }
}
