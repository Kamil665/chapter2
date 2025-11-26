#include <stdio.h>

int main() {
    int num = 634;
    
    int n1 = num / 100;
    int n2 = (num % 100) / 10;
    int n3 = (num % 100) % 10;
    
    printf("Трёхзначное число: %d", num);
    printf("\n");
    printf("\n");
    printf("%d%d%d", n2, n1, n3);
    printf("\n");
    printf("%d%d%d", n2, n3, n1);
    printf("\n");
    printf("%d%d%d", n1, n3, n2);
    printf("\n");
    printf("%d%d%d", n3, n1, n2);
    printf("\n");
    printf("%d%d%d", n3, n2, n1);
    printf("\n");
    printf("%d%d%d", n1, n2, n3);
}
