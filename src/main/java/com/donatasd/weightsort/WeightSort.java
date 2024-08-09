package com.donatasd.weightsort;

import java.util.Arrays;

public class WeightSort {

    public static String orderWeight(String value) {
        var list = Arrays.stream(value.split(" ")).map(Weight::new).sorted().map(Weight::toString).toList();
        StringBuilder sbStr = new StringBuilder();
        for (int i = 0; i < list.size(); i++) {
            sbStr.append(list.get(i));
            if (list.size() - 1 != i) {
                sbStr.append(" ");
            }
        }
        return sbStr.toString();
    }

    public static class Weight implements Comparable<Weight> {
        private String weight;
        private Integer value;

        public Weight(String weight) {
            this.weight = weight;
            this.value = weight.chars().mapToObj(Character::getNumericValue).reduce(0, Integer::sum);
        }

        @Override
        public int compareTo(Weight o) {
            var result = this.value.compareTo(o.value);
            return result == 0 ? this.weight.compareTo(o.weight) : result;
        }

        @Override
        public String toString() {
            return this.weight;
        }
    }
}