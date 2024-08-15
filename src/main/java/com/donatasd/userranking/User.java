package com.donatasd.userranking;

public class User {

    public int rank;
    public int progress;

    public User() {
        this.rank = -8;
        this.progress = 0;
    }

    public void incProgress(int rank) {
        if (rank == 0 || rank > 8 || rank < -8) {
            throw new IllegalArgumentException();
        }
        if (this.rank != 8) {
            var totalProgress = calcProgress(rank) + this.progress;
            this.rank += totalProgress / 100;
            this.progress = this.rank == 8 ? 0 : totalProgress % 100;
            if (this.rank == 0) {
                this.rank++;
            }
        }
    }

    private int calcProgress(int rank) {
        if (rank == this.rank) {
            return 3;
        }
        if (rank < this.rank) {
            return 1;
        }
        var diff = Math.abs(rank - this.rank);
        if (rank > 0 && this.rank < 0) {
            diff -= 1;
        }
        return 10 * diff * diff;

    }
}
