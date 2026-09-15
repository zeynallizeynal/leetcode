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

  removeFirst() {
    const val = this.arr[this.start];
    this.arr[this.start] = null;
    this.start = (this.start + 1) % this.size;
    this.count--;
    return val;
  }

  addLast(v) {
    this.arr[this.end] = v;
    this.end = (this.end + 1) % this.size;
    this.count++;
  }

  removeLast() {
    this.end = (this.end - 1 + this.size) % this.size;
    const val = this.arr[this.end];
    this.arr[this.end] = null;
    this.count--;
    return val;
  }

  get first() {
    return this.arr[this.start];
  }

  get last() {
    return this.arr[(this.end - 1 + this.size) % this.size];
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
ca.addLast(43);
ca.addLast(56);
console.log(ca.start, ca.end);
console.log(ca.first, ca.last);

console.log(ca.removeLast());
console.log(ca.first, ca.last);
