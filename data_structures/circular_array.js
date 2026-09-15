class CircularArray {
  constructor(size = 1) {
    this.size = size;
    this.arr = new Array(this.size);

    this.count = 0;
    this.start = 0;
    this.end = 0;
  }

  addFirst(v) {
    this.start = (this.start - 1 + this.size) % this.size;
    this.arr[this.start] = v;
    this.count++;
  }

  addLast(v) {
    this.arr[this.end] = v;
    this.end = (this.end + 1) % this.size;
    this.count++;
  }

  get first() {
    return this.arr[this.start];
  }

  get last() {
    return this.arr[this.end];
  }

  // print() {
  //   let count = 0;
  //   for (let i = 0; i < this.arr.length; i = (i + 1) % this.arr.length) {
  //     console.log(this.arr[i]);
  //     if (++count >= 20) return;
  //   }
  // }
}

// [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
const ca = new CircularArray(10);
ca.addFirst(95);
console.log(ca.start);