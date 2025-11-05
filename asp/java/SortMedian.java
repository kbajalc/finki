public class SortMedian {
    int Delay;
    int Length;
    boolean Strict;

    int Sum;
    int Last;
    int Value;
    int First;
    int Pivot;
    int Index;
    boolean Ready;
    int[] Buffer;
    int[] Sorted;

    public SortMedian(int length) {
        this.Delay = length / 2;
        this.Length = length;
        this.Strict = true;
        this.Buffer = new int[length];
        this.Sorted = new int[length];

        for (int i = 0; i < length; i++) {
            this.Buffer[i] = 0;
            this.Sorted[i] = 0;
        }
    }

    public void Reset() {
        this.Sum = 0;
        this.Last = 0;
        this.Value = 0;
        this.First = 0;
        this.Pivot = 0;
        this.Index = 0;
        this.Ready = false;
        for (int i = 0; i < this.Length; i++) {
            this.Buffer[i] = 0;
            this.Sorted[i] = 0;
        }
    }

    public int Filter(int datum) {
        if (!this.Ready && this.Index == 0) {
            for (int i = 0; i < this.Length; i++) {
                this.Buffer[i] = datum;
                this.Sorted[i] = datum;
                this.Sum += datum;
            }
        }

        int rem = this.Buffer[this.Index];
        this.Buffer[this.Index] = datum;
        this.Index++;
        if (this.Index == this.Length) {
            this.Index = 0;
            this.Ready = true;
        }
        this.First = this.Buffer[this.Index];

        this.Sum += datum - rem;
        this.Last = datum;

        int half = this.Index + this.Delay;
        if (half >= this.Length) {
            half -= this.Length;
        }
        this.Pivot = this.Buffer[half];

        int ix = 0;
        for (int i = 0; i < this.Length; i++) {
            if (this.Sorted[i] == rem) {
                ix = i;
                break;
            }
        }

        this.Sorted[ix] = datum;
        for (int i = ix; i < this.Length - 1; i++) {
            if (this.Sorted[i] > this.Sorted[i + 1]) {
                int temp = this.Sorted[i];
                this.Sorted[i] = this.Sorted[i + 1];
                this.Sorted[i + 1] = temp;
            } else {
                break;
            }
        }
        for (int i = ix; i > 0; i--) {
            if (this.Sorted[i] < this.Sorted[i - 1]) {
                int temp = this.Sorted[i];
                this.Sorted[i] = this.Sorted[i - 1];
                this.Sorted[i - 1] = temp;
            } else {
                break;
            }
        }

        int m = this.Delay;
        if (this.Strict) {
            this.Value = this.Sorted[m];
        } else {
            this.Value = this.Sorted[m - 1] + this.Sorted[m] + this.Sorted[m + 1] / 3;
        }

        return this.Value;
    }
}
