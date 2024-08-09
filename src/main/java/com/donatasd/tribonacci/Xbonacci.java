package com.donatasd.tribonacci;


public class Xbonacci {

    public double[] tribonacci(double[] s, int n) {
        var result = new double[n];
        for (int i = 0; i < n; i++) {
            if (i < s.length) {
                result[i] = s[i];
            } else {
                result[i] = result[i - 3] + result[i - 2] + result[i - 1];
            }
        }
        return result;
    }
}