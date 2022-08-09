# Calculate Shipping CLI

This is a CLI that calculates the total delivery cost of each package in a delivery vehicle,
and applies an entered offer code per package, if the coupon is valid and satisfies conditions.

Offer codes can be found in cmd/offercodes.json

DeliveryCost = BaseDeliveryCost + PackageToTalWeight * 10 + DistanceToDestination * 5

## Input:

Base Delivery Cost<br>
Number of Packages

For each package:
    <ul>
    <li>Package ID</li>
    <li>Package Weight in Kg</li>
    <li>Distance in Km</li>
    <li>Offer Code</li>
    </ul>
    
## Output:

For each package:
    <ul>
    <li>Package ID</li>
    <li>Discount Applied</li>
    <li>Total Cost</li>
</ul>

## Offer sample

 <table>
  <tr>
    <th>Offer Code<br>Discount</th>
    <th>Distance (Km)</th>
      <th>Weight (Kg)</th>
  </tr>
  <tr>
    <td>OFR001<br>10% Discount</td>
    <td>< 200</td>
    <td>70 - 200</td>
  </tr>
      <tr>
    <td>OFR002<br>7% Discount</td>
    <td>50 - 150</td>
    <td>100 - 250</td>
  </tr>
      <tr>
    <td>OFR003<br>5% Discount</td>
    <td>50 - 250</td>
    <td>10 - 150</td>
  </tr>
</table> 
