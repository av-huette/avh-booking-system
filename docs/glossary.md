# Glossary

The terms used in AVHBS, in its code and in its documentation.
Several of them mean something narrower than in everyday language; **account** and **user** in particular are different things.

- **Account:** A member's account at the bar.
  It has a balance and a maximum debt, belongs to an account category and can be disabled; it is never deleted ([ADR 0004](adr/0004-traceable-bookings.md)).
- **Balance:** The money on an account, in cents ([ADR 0005](adr/0005-money-as-integer-cents.md)).
  It changes only together with a recorded order or payment ([ADR 0004](adr/0004-traceable-bookings.md)).
- **Booking:** Several orders booked together.
- **Category:** A way to group accounts or products.
  A category is either an account category or a product category.
- **Client:** One connected till or device, identified by its client ID.
  Not a customer.
- **Favorite:** A product that an account orders often, with a count of how often.
- **GoBD:** The German principles for proper bookkeeping in electronic form.
  Among other things, nothing booked may disappear or change without a trace.
- **KassenSichV:** The German regulation on the security of electronic cash registers; it requires a TSE.
- **Location:** The physical place where a till stands.
- **Maximum debt:** How far an account's balance may go below zero.
- **Order:** A purchase by one account, made up of order lines.
  Each order line holds a product, the amount and the product's price at the time of the order.
- **Payment:** Money paid onto an account: a top-up, or the settling of a debt.
- **Product:** Something that can be bought at the bar, with a price, a VAT rate, a size and unit, a product group and a product category.
- **Product group:** A group of products in the till's menu.
  Product groups can be nested.
- **Product visibility:** A rule that decides which products a till shows, depending on its location and the account category.
- **Rights:** A permission that a role grants or denies.
- **Role:** A set of rights that a user has.
- **Sewobe:** External software for managing the members of an association ([sewobe.de](https://www.sewobe.de/)).
  An association that uses it can link its members to AVHBS accounts; AVHBS works without it.
- **Settings:** Key-value configuration, kept in three groups: frontend, payment and email.
- **TSE:** Technical security device (technische Sicherheitseinrichtung): the certified component that signs every transaction under the KassenSichV.
- **Unit:** The unit of a product's size, for example liter or piece.
- **User:** Someone who logs in to manage AVHBS, with a role.
  Not the same as an account: a user manages the system, an account holds a member's balance.
- **VAT:** A value added tax rate, stored as a percentage (for example 19).
