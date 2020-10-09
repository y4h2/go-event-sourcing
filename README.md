# go-event-sourcing

go event sourcing example

## Plan

### v1

Account service relies on event store directly.
All data stores in memory.

## Main Logic

- Open: open an account with init balance
- Transfer: transfer money from one account to another account
- Store: store money to the account
- Withdraw: withdraw money from the account

## Problem

Everytime users get balance, they need to get all event list and do aggregation.

Transfer, store, and withdraw don't verify whether account exist.

## Following Work

Create a snapshot store to store aggregation result.

Create an entity store to store exist account.
