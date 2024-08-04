package com.donatasd.tribonacci;


import org.junit.jupiter.api.*;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

public class XbonacciTest {
    private Xbonacci variabonacci;

    @BeforeEach
    public void setUp() throws Exception {
        variabonacci = new Xbonacci();
    }

    @AfterEach
    public void tearDown() throws Exception {
        variabonacci = null;
    }

    private final double precision = 1e-10;

    @Test
    public void sampleTests() {
        assertArrayEquals(new double []{1,1,1,3,5,9,17,31,57,105}, variabonacci.tribonacci(new double []{1,1,1},10), precision);
        assertArrayEquals(new double []{0,0,1,1,2,4,7,13,24,44}, variabonacci.tribonacci(new double []{0,0,1},10), precision);
        assertArrayEquals(new double []{0,1,1,2,4,7,13,24,44,81}, variabonacci.tribonacci(new double []{0,1,1},10), precision);
    }
}
