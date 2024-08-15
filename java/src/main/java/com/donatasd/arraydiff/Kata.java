package com.donatasd.arraydiff;

import java.util.Arrays;

public class Kata {

    public static int[] arrayDiff(int[] a, int[] b) {
        var bArray = Arrays.stream(b).boxed().toList();
        return Arrays.stream(a).filter((value) -> !bArray.contains(value)).toArray();
    }

}