
#include <stdlib.h>

int findClosest(int x, int y, int z) {
    int diff1 = abs(x - z);
    int diff2 = abs(y - z);

    return (diff1 < diff2) ? 1 : (diff2 < diff1) ? 2 : 0;
}

