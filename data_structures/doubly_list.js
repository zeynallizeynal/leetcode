function ListNode(v = null) {
  this.val = v;
  this.next = null;
  this.prev = null;
}

function createList(arr) {
  if (!arr || arr.length === 0) return null;

  const head = new ListNode(arr[0]);
  let cur = head;

  for (let i = 1; i < arr.length; i++) {
    const newNode = new ListNode(arr[i]);
    cur.next = newNode;
    newNode.prev = cur;
    cur = cur.next;
  }

  return head;
}

function traverseList(head) {
  for (let temp = head; temp !== null; temp = temp.next) {
    console.log(temp.val);
  }
}

function traverseFromTail(head) {
  for (let temp = head; temp !== null; temp = temp.prev) {
    console.log(temp.val);
  }
}

function insertTail(head, v) {
  let temp = head;
  while (temp.next !== null) {
    temp = temp.next;
  }
  const newNode = new ListNode(v);
  temp.next = newNode;
  newNode.prev = temp;
  temp = newNode;
  return temp
}

function insertHead(head, v) {
  const newNode = new ListNode(v);
  newNode.next = head;
  head.prev = newNode;
  head = newNode;

  return head;
}

function deleteHead(head) {
  head = head.next;
  return head;
}

function deleteTail(head) {
  let temp = head;
  while (temp.next.next !== null) {
    temp = temp.next;
  }
  temp.next = null;
}

let listHead = createList([1, 2, 3, 4, 5]);
insertTail(listHead, 28);
listHead = insertHead(listHead, 9);
// listHead = deleteHead(listHead);
// deleteTail(listHead);
traverseFromTail(listHead);