import numpy as np


def shuffel(deck):
    left_deck = deck[:len(deck)/2].copy()
    right_deck = deck[len(deck)/2:].copy()

    for i in range(len(left_deck)):
        deck[i * 2] = left_deck[i]
        deck[i * 2 + 1] = right_deck[i]
    return deck

def compare(a,b):
    for i in range(len(a)):
        if a[i] != b[i]:
            return False
    return True

def main():
    number_of_cards = 50000

    deck = np.arange(1, number_of_cards+1)
    orginal = deck.copy()

    i = 1

    print orginal
    print deck
    shuffel(deck)

    while compare(orginal, deck) == False:
        i += 1
        shuffel(deck)

    print i

main()