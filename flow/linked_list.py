from __future__ import annotations
from dataclasses import dataclass


@dataclass(frozen=True)
class Node[T]:
    value: T
    ptr: None | Node[T]


@dataclass(frozen=True)
class LinkedList[T]:
    ptr: None | Node[T]

    @classmethod
    def empty(cls) -> LinkedList[T]:
        return LinkedList(None)

    @classmethod
    def singleton(cls, value: T) -> LinkedList[T]:
        return LinkedList(Node(value , None))

    def cons(self, value: T) -> LinkedList[T]:
        return LinkedList(Node(value, self.ptr))

    def head(self) -> Node[T]:
        if self.ptr is None:
            raise ValueError('head of empty list')
        return self.ptr

    def tail(self) -> LinkedList[T]:
        return LinkedList(self.ptr.ptr if self.ptr is not None else None)

    def null(self) -> bool:
        return self.ptr is None
