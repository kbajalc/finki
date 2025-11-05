import java.util.Random;

public class MedianTest {
    private static void TestMedianRandom() {
        SortMedian f = new SortMedian(11);
        HeapMedian m = new HeapMedian(11);

        Random rng = new Random(42);
        for (int i = 0; i < 1 * 3600 * 250; i++) {
            int v = i + rng.nextInt(1000);
            int a = f.Filter(v);
            int b = m.Filter(v);
            if (m.Ready && a != b) {
                System.out.println(">>> MISMATCH " + i + " " + v + " " + a + " " + b);
                throw new IllegalStateException("Median mismatch");
            }
        }
        System.out.println("Median random test passed");
    }

    private static void TestMedianPerformance() {
        int WIN = (600 / 2) | 1;
        int LEN = 10 * 3600 * 250;
        int[] DATA = new int[LEN];
        {
            Random rng = new Random(42);
            for (int i = 0; i < LEN; i++) {
                DATA[i] = rng.nextInt(1000);
            }
        }

        long srt;
        {
            SortMedian med = new SortMedian(WIN);
            long start = System.nanoTime();
            for (int i = 0; i < LEN; i++) {
                int v = DATA[i];
                med.Filter(v);
            }
            srt = (System.nanoTime() - start) / 1000;
            System.out.println("SortMedian took: " + srt);
        }

        long hep;
        {
            HeapMedian med = new HeapMedian(WIN);
            long start = System.nanoTime();
            for (int i = 0; i < LEN; i++) {
                int v = DATA[i];
                med.Filter(v);
            }
            hep = (System.nanoTime() - start) / 1000;
            System.out.println("HeapMedian took: " + hep);
        }

        System.out.printf("Heap/Sort = %.3f%%\n", (double)hep * 100 / (double)srt - 100);
    }

    public static void main(String[] args) {
        TestMedianRandom();
        TestMedianPerformance();
    }
}
