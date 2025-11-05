public class HeapMedian {
    int Length;
	int Delay;
	int Index;
	boolean Ready;
	int First;
	int Pivot;

	int Value;

	int[] Data;
	Heap Small; // lower half (extra element for odd)
	Heap Large; // upper half

    public HeapMedian(int length) {
        if (length <= 0) {
            throw new IllegalArgumentException("window size must be positive");
        }
        this.Length = length;
		this.Delay = length / 2;
		this.Value = 0;
		this.Data = new int[length];
		this.Small = new Heap(length / 2 + 2, false);
		this.Large = new Heap(length / 2 + 2, true);
    }

    public void Reset() {
        this.Index = 0;
        this.Ready = false;
        this.First = 0;
        this.Pivot = 0;

        this.Value = 0;
        for (int i = 0; i < this.Data.length; i++) {
            this.Data[i] = 0;
        }
        this.Small.Reset();
        this.Large.Reset();
    }

    public int Filter(int x) {
        if (this.Ready) {
            this.Pop(this.Index);
        }
        this.Push(this.Index, x);

        this.Index++;
        if (this.Index == this.Length) {
            this.Ready = true;
            this.Index = 0;
        }
        this.First = this.Data[this.Index];

        var half = this.Index + this.Delay;
        if (half >= this.Length) {
            half -= this.Length;
        }
        this.Pivot = this.Data[half];

        int val = 0;
        if (this.Small.Len > this.Large.Len) {
            val = this.Small.Top();
        } else if (this.Small.Len == this.Large.Len) {
            val = (this.Small.Top() + this.Large.Top()) / 2;
        } else {
            val = this.Large.Top();
        }
        this.Value = val;
        return val;
    }

    public void Pop(int head) {
        var val = this.Data[head];
        this.remove(val);
        this.rebalance();
    }

    public void Push(int tail, int x) {
        this.Data[tail] = x;
        if (this.Small.Len == 0 || (this.Small.Len > 0 && x <= this.Small.Top())) {
            this.Small.Push(x);
        } else {
            this.Large.Push(x);
        }
        this.rebalance();
    }

    private int targetSizes() {
        var count = this.Small.Len + this.Large.Len;
        if (count == 0) {
            return 0;
        }
        if ((count & 1) == 1) {
            return (count + 1) / 2;
        } else {
            return count / 2;
        }
    }

    private void rebalance() {
        int wantS = this.targetSizes();
        while (this.Small.Len > wantS) {
            int v = this.Small.Pop();
            this.Large.Push(v);
        }
        while (this.Large.Len > 0 && this.Small.Len < wantS) {
            int v = this.Large.Pop();
            this.Small.Push(v);
        }
    }

    private void remove(int val) {
        if (val <= this.Small.Top()) {
            if (this.Small.remove(val)) {
                return;
            }
        }
        if (val >= this.Large.Top()) {
            if (this.Large.remove(val)) {
                return;
            }
        }
    }
}
