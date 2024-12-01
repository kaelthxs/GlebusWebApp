import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        int[][] arrayOfBakers = inputInArray();
        int[] Concurents = checkCountOfConcurrents(arrayOfBakers);
        for (int i = 0; i < Concurents.length; i++) {
            System.out.print(Concurents[i] + " ");
        }
    }

    public static int[][] inputInArray() {
        Scanner input = new Scanner(System.in);
        int N = input.nextInt();
        int[][] arrayOfBakers = new int[N][4];
        for (int i = 0; i < arrayOfBakers.length; i++) {
            for (int j = 0; j < arrayOfBakers[i].length; j++) {
                arrayOfBakers[i][j] = input.nextInt();
            }
        }
        return arrayOfBakers;
    }

    public static int[] checkCountOfConcurrents(int[][] arrayOfBakers) {
        int[] Concurents = new int[arrayOfBakers.length];
        for (int i = 0; i < arrayOfBakers.length; i++) {
            for (int j = 0; j < arrayOfBakers.length; j++) {
                if (i != j && checkConcurrency(arrayOfBakers[i], arrayOfBakers[j])) {
                    Concurents[i]++;
                }
            }
        }
        return Concurents;
    }

    public static boolean checkConcurrency(int[] arrayOfBakers1, int[] arrayOfBakers2) {
        if ((arrayOfBakers1[2] <= arrayOfBakers2[0]) || (arrayOfBakers1[0] >= arrayOfBakers2[2]) || (arrayOfBakers1[3] <= arrayOfBakers2[1]) || (arrayOfBakers1[1] >= arrayOfBakers2[3])){
            return false;
    }
        return true;
}

}
