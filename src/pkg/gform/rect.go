package gform

import (
    "w32"
    "w32/user32"
)

type Rect struct {
    rect w32.RECT
}

func NewEmptyRect() *Rect {
    var newRect Rect
    user32.SetRectEmpty(&newRect.rect)

    return &newRect
}

func NewRect(left, top, right, bottom int) *Rect {
    var newRect Rect
    user32.SetRectEmpty(&newRect.rect)
    newRect.Set(left, top, right, bottom)

    return &newRect
}

func (this *Rect) Get() (left, top, right, bottom int) {
    left = this.rect.Left
    top = this.rect.Top
    right = this.rect.Right
    bottom = this.rect.Bottom
    return
}

func (this *Rect) GetW32Rect() *w32.RECT {
    return &this.rect
}

func (this *Rect) Set(left, top, right, bottom int) {
    user32.SetRect(&this.rect, left, top, right, bottom)
}

func (this *Rect) IsEqual(rect *Rect) bool {
    return user32.EqualRect(&this.rect, &rect.rect)
}

func (this *Rect) Inflate(x, y int) {
    user32.InflateRect(&this.rect, x, y)
}

func (this *Rect) Intersect(src *Rect) {
    user32.IntersectRect(&this.rect, &this.rect, &src.rect)
}

func (this *Rect) IsEmpty() bool {
    return user32.IsRectEmpty(&this.rect)
}

func (this *Rect) Offset(x, y int) {
    user32.OffsetRect(&this.rect, x, y)
}

func (this *Rect) IsPointIn(x, y int) bool {
    return user32.PtInRect(&this.rect, x, y)
}

func (this *Rect) Substract(src *Rect) {
    user32.SubtractRect(&this.rect, &this.rect, &src.rect)
}

func (this *Rect) Union(src *Rect) {
    user32.UnionRect(&this.rect, &this.rect, &src.rect)
}
