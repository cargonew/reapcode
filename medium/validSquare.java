import java.util.*;

class Solution {
    public boolean validSquare(int[] p1, int[] p2, int[] p3, int[] p4) {
        int[] dists = new int[] {
            distSq(p1, p2), distSq(p1, p3), distSq(p1, p4),
            distSq(p2, p3), distSq(p2, p4), distSq(p3, p4)
        };
        
        Arrays.sort(dists);

        if (dists[0] == 0) return false;

        boolean sidesEqual = dists[0] == dists[1] &&
                             dists[1] == dists[2] &&
                             dists[2] == dists[3];

        boolean diagsEqual = dists[4] == dists[5] &&
                             dists[4] > dists[3];

        return sidesEqual && diagsEqual;
    }

    public int distSq(int[] a, int[] b) {
        int dx = a[0] - b[0];
        int dy = a[1] - b[1];
        return dx * dx + dy * dy;
    }
}

