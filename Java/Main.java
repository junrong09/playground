import java.io.*;

public class Main {
  public static void main(String[] args) {
    try {
      //BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
      BufferedReader br = new BufferedReader(new FileReader("a.txt"));
      String s;
      while ((s = br.readLine()) != null) {
        System.out.println(s);
      }
    } catch (Exception e) {

    }
  }
}
