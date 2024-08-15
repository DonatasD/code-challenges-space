package com.donatasd.oddoreven;

import java.util.Arrays;

public class Codewars {
    public static String oddOrEven (int[] array) {
        return Arrays.stream(array).asLongStream().reduce(0, (acc, cur) -> (acc + cur) % 2) % 2 == 0 ? "even" : "odd";
    }
}