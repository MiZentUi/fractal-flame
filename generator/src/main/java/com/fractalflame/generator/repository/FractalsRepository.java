package com.fractalflame.generator.repository;

import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.fractalflame.generator.entity.Fractal;

@Repository
public interface FractalsRepository extends JpaRepository<Fractal, Long> {
    Page<Fractal> findAllByUserId(Long id, Pageable pageable);
}
