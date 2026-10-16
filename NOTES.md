# Notes

## Assumptions and decisions

- The group creator is the first member and receives the first payout.
- Members are added in join order and receive the pot in that order.
- A member can belong to only one group.
- A member can be added after a week has started. They must pay in the current week like every other member.
- A week cannot close until every member has paid the full weekly amount.
- A partial payment remains recorded until the member finishes paying.
- The program manages one active group at a time through the current menu flow, although the data maps can store more than one group.
- The program currently uses two-decimal `float64` values for money to keep the implementation simple.

## Improvements with more time

- Store money as integer pesewas instead of `float64`.
- Store the week number with every payment.
- Keep payment data and payout history in clearer group structures.
- Add an option to exit the program.
- Add a clearer rule for adding members after the first week has closed.
- Make the group report and payout history persistent.
