function corpFlightBookings(bookings: number[][], n: number): number[] {
  let diff: number[] = new Array(n).fill(0);
  let seats: number[] = new Array(n).fill(0);
  bookings.forEach(el => {
    let beginIndex = el[0] - 1;
    let endIndex = el[1] - 1;
    let addNumber = el[2];
    diff[beginIndex] += addNumber;
    diff[endIndex + 1] -= addNumber;
  });

  seats.forEach((_, index) => {
    if (index === 0) {
      seats[0] = diff[0];
    } else {
      seats[index] = diff[index] + seats[index - 1];
    }
  });

  return seats
}
