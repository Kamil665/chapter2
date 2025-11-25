public class MyClass {
  public static void main(String args[]) {
      int num = 236;
      
      int x1 = num / 100;
      int x2 = (num % 100) / 10;
      int x3 = (num % 100) % 10;
      
      System.out.println("Трёхзначное число: " + num);
      System.out.println("После исправления: " + x3 + x1 + x2);
  }
}
