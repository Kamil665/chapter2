#import <Foundation/Foundation.h>

int main(int argc, const char * argv[]) {
    @autoreleasepool {
        int num = 5678;
        
        int n1 = num / 1000;
        int n2 = (num % 1000) / 100;
        int n3 = ((num % 1000) % 100) / 10;
        int n4 = ((num % 1000) % 100) % 10;
        
        int sum = n1 + n2 + n3 + n4;
        int multi = n1 * n2 * n3 * n4;
        
        NSLog(@"Четырёхзначное число: %d", num);
        NSLog(@"Сумма его цифр: %d", sum);
        NSLog(@"Произведение его цифр: %d", multi);
        }
    return 0;
}
