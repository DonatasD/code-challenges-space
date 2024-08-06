package com.donatasd.permutations;

import org.junit.jupiter.api.Test;

import java.util.*;
import java.util.stream.*;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class SolutionTest {

    @Test
    public void example1() {
        assertEquals( new ArrayList<String>(Arrays.asList("a")),
                Permutations.singlePermutations("a").stream().sorted().collect(Collectors.toList()) );
    }

    @Test public void example2() {
        assertEquals( new ArrayList<String>(Arrays.asList("ab","ba")),
                Permutations.singlePermutations("ab").stream().sorted().collect(Collectors.toList()) );
    }

    @Test public void example3() {
        assertEquals( new ArrayList<String>(Arrays.asList("aabb", "abab", "abba", "baab", "baba", "bbaa")),
                Permutations.singlePermutations("aabb").stream().sorted().collect(Collectors.toList()) );
    }
}