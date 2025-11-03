import type { Party } from 'gen/guest/v1/guest_pb';

export function formatNames(party: Party): string {
  let result = '';
  for (const [i, guest] of party.guests.entries()) {
    if (i > 0) {
      result += i === party.guests.length - 1 ? ' & ' : ', ';
    }
    result += guest.firstName;
  }
  return result;
}
