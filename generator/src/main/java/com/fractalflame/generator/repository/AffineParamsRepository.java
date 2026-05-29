package com.fractalflame.generator.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.fractalflame.generator.entity.AffineParams;

@Repository
public interface AffineParamsRepository extends JpaRepository<AffineParams, Long> {
}
