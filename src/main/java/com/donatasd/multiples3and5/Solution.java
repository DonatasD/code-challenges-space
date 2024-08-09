package com.donatasd.multiples3and5;

public class Solution {

    public int solution(int number) {
        var lastNumber = number - 1;
        var threesSum = calcArithmeticSeqSum(3, 3, lastNumber / 3);
        var fivesSum = calcArithmeticSeqSum(5, 5, lastNumber / 5);
        var fifteenSum = calcArithmeticSeqSum(15, 15, lastNumber / 15);
        return threesSum + fivesSum - fifteenSum;
    }

    private int calcArithmeticSeqSum(int firstTerm, int commonDiff, int totalNumbersInSeq) {
        return (totalNumbersInSeq * ((2 * firstTerm) + (totalNumbersInSeq - 1) * commonDiff)) / 2;
    }
}
