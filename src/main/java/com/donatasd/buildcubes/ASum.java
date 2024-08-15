package com.donatasd.buildcubes;

public class ASum {

    public static long findNb(long m) {
        var result = 0;
        var sum = 0L;

        for (int i = 1; result == 0; i++) {
            sum += (long) Math.pow(i, 3);
            if (m == sum) {
                result = i;
            }
            if (sum > m) {
                result = -1;
            }
        }
        return result;
    }
}