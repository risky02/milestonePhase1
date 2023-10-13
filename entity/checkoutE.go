package entity

import "time"

type Checkout struct {
    CheckoutID   int
    UserID       int
    TotalAmount  float64
    CheckoutDate time.Time
}