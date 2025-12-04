#include <iostream>
using namespace std;

int main() {
    int n;
    cout << "Введите число n (10 <= n <= 999): ";
    cin >> n;
    cout << "\n";
    
    if ((n < 10) || (n > 999)) {
        cout << "Ошибка! Диапазон нарушен!" << endl;
    }
    else if (n % 10 == 0) {
        cout << "Ошибка! Десяток числа не должен быть равен нулю!" << endl;
    }
    else {
        cout << "n = " << n << endl;
        
        int num1 = n / 100;
        int num2 = (n % 100) * 10;
        int x = num1 + num2;
        
        cout << "x = " << x << endl;
    }
}
