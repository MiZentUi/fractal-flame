package com.fractalflame.generator.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.fractalflame.generator.entity.Function;

@Repository
public interface FunctionsRepository extends JpaRepository<Function, Long> {
}
