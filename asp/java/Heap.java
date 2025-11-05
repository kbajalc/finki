public class Heap {
    int Cap;   // maximum number of elements
    boolean Min;  // max-heap if false, min-heap if true
    int Len;   // current number of elements
    int[] Buf; // backing array, length = capacity

    public Heap(int capy, boolean mini) {
        this.Cap = capy;
        this.Min = mini;
        this.Len = 0;
        this.Buf = new int[capy];
        for (int i = 0; i < capy; i++) {
            this.Buf[i] = 0;
        }
    }

    public void Reset() {
        this.Len = 0;
        for (int i = 0; i < this.Cap; i++) {
            this.Buf[i] = 0;
        }
    }   

    public int Top() {
        if (this.Len == 0) {
            throw new IllegalStateException("Heap.Top on empty heap");
        }
        return this.Buf[0];
    }

    public void Push(int x) {
        if (this.Len == this.Buf.length) {
            throw new IllegalStateException("Heap overflow");
        }
        this.Buf[this.Len] = x;
        this.up(this.Len);
        this.Len++;
    }

    public int Pop() {
        if (this.Len == 0) {
            throw new IllegalStateException("Heap.Pop on empty heap");
        }
        int top = this.Buf[0];
        this.Len--;
        this.swap(0, this.Len);
        this.down(0);
        return top;
    }

    private boolean cmp(int i, int j) {
        int ai = this.Buf[i];
        int aj = this.Buf[j];
        if (this.Min) {
            return ai < aj;
        } else {
            return ai > aj;
	    }
    }

    private void swap(int i, int j) {
        int temp = this.Buf[i];
        this.Buf[i] = this.Buf[j];
        this.Buf[j] = temp;
    }

    private void up(int i) {
        while (true) {
            if (i == 0) {
                return;
            }
            int p = (i - 1) >> 1;
            if (!this.cmp(i, p)) {
                return;
            }
            this.swap(i, p);
            i = p;
        }
    }

    private void down(int i) {
        while (true) {
            int l = i*2 + 1;
            if (l >= this.Len) {
                return;
            }
            int j = l;
            int r = l + 1;
            if (r < this.Len && this.cmp(r, l)) {
                j = r;
            }
            if (!this.cmp(j, i)) {
                return;
            }
            this.swap(i, j);
            i = j;
        }
    }

    // Remove the deepest occurrence of val from the heap.
    protected boolean remove(int val) {
        for (int i = this.Len - 1; i >= 0; i--) {
            if (this.Buf[i] == val) {
                this.removeAt(i);
                return true;
            }
        }
        return false;
    }

    // Remove the element at index i from the heap.
    private int removeAt(int i) {
        if (i == 0) {
            return this.Pop();
        }
        if (i >= this.Len) {
            throw new IllegalArgumentException("invalid index");
        }
        int r = this.Buf[i];
        this.Len--;
        this.swap(i, this.Len);
        // Try both up and down since the swapped element
        // may violate heap property in either direction.
        this.up(i);
        this.down(i);
        return r;
    }
}   
