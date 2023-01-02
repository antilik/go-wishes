package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func getRandomNumber(number int) (int, error) {
	if number < 0 {
		return 0, fmt.Errorf("Passed number cannot be negative")
	}
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(number) + 1
	return target, nil
}
func getName() (string, error) {
	fmt.Print("Enter your name and press 'Enter': ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	return input, nil
}

func main() {
	customerFirstName, err := getName()
	if err != nil {
		log.Fatal(err)
	}

	sliceWishesInspirational := []string{
		"A New Year is like a blank book, and the pen is in your hands. It is your chance to write a beautiful story for yourself. Happy New Year.",
		"As the New Year approaches us with hopes anew, here is to wishing you and your family a wonderful year ahead.",
		"As the New Year dawns, I hope it is filled with the promises of a brighter tomorrow. Happy New Year!",
		"Every end marks a new beginning. Keep your spirits and determination unshaken, and you shall always walk the glory road. With courage, faith and great effort, you shall achieve everything you desire. I wish you a Happy New Year.",
		"On the road to success, the rule is always to look ahead. May you reach your destination and may your journey be wonderful. Happy New Year.",
		"No one can go back in time to change what has happened, so work on your present to make yourself a wonderful future.",
		"You’re supposed to let go of the past and start off new. You’re supposed to forgive all those who have hurt you and be open to new relationships, with open arms. That is why it is called the ‘New’ Year. May you have a Happy New Year.",
		"End each year with a few good lessons and start the new one by showing that you have learnt the lessons of the past well.",
		"The New Year has brought another chance for us to set things right and to open up a new chapter in our lives.",
		"Unlike what most people think it is never too late to be what you wish to be. New Years card with religious message",
		"Failure doesn’t influence your inner resilience, and failing simply means that you’ve discovered another false way to move your life forward. Profit from it.",
		"At precisely the moment you feel like giving up, victory is always within reach. Remember this moving into the New Year.",
		"We will open the book. Its pages are blank. We are going to put words on them ourselves. The book is called Opportunity, and its first chapter is New Year’s Day.",
	}
	sliceWishesClassic := []string{
		"Wishing you a Happy New Year with the hope that you will have many blessings in the year to come.",
		"Out with the old, in with the new: may you be happy the whole year through. Happy New Year!",
		"Counting my blessings and wishing you more. I hope you enjoy the New Year in store.",
		"I resolve to stop wasting my resolutions on myself and use them to repay you for the warmth you’ve shown me. Happy New Year!",
		"Nights will be dark but days will be light, wishing your life to be always bright – Happy New Year.",
		"Let us look back at the past year with the warmest of memories. Happy New Year.",
		"Let the old year end and the New Year begin with the warmest of aspirations. Happy New Year!",
		"One more year loaded with sweet recollections and cheerful times has passed. You have made my year exceptionally uncommon, and I wish this continues forever. With you around, each minute is a unique event for me. I wish you to have a year as incredible as you are.",
		"On this New Year, I wish that you have a superb January, a dazzling February, a Peaceful March, an anxiety-free April, a sensational May, and joy that keeps going from June to November, and then round off with an upbeat December.",
		"On this New Year, may you change your direction and not dates, change your commitments and not the calendar, change your attitude and not the actions, and bring about a change in your faith, your force, and your focus and not the fruit. May you live up to the promises you have made and may you create for you and your loved ones the happiest New Year ever.",
		"May this year bring new happiness, new goals, new achievements, and a lot of new inspirations on your life. Wishing you a year fully loaded with happiness.",
		"Wishing every day of the new year to be filled with success, happiness, and prosperity for you. Happy New Year.",
		"May the new year bring you warmth, love, and light to guide your path to a positive destination.",
		"Here’s wishing you all the joy of the season. Have a Happy New Year!",
	}
	sliceWishesFunny := []string{
		"Wishing you 12 months of success, 52 weeks of laughter, 365 days of fun, 8,760 hours of joy, 525,600 minutes of good luck, and 31,536,000 seconds of happiness.",
		"Many people look forward to the New Year for a new start on old habits.",
		"I’ve been waiting 365 days to say “Happy New Year” since I had so much fun saying it last year. Happy New Year, Friend.",
		"May this New Year bring actual change in you, not a recurrence of old habits in a new package.",
		"A New Year is the chance to start over with a clean slate. Too bad my credit card won’t start over with a blank slate.",
		"The most fun part about making New Year’s resolutions is breaking New Year’s resolutions. Can’t wait to mess up with you.",
		"I’m so excited for the new year. Unfortunately, I don’t have any resolutions to make since I’m already perfect.",
		"Happy New Year. Here is a wish for the New Year from someone who is adorable, handsome, and intelligent and wants to see you smiling always.",
		"Happy New Year. Here’s to another year pretending that I like you people.",
		"Happy New Year. Here’s to having a fresh start at binge eating, boozing, and slacking off.",
		"We all get the exact same 365 days. The only difference is what we do with them.",
		"This year may I wish you finally learn how to use your smartphone properly.",
		"To a New Year full of new possibilities, even though I’m sure we’ll just do the same old stuff anyway.",
		"Happy New Year and good luck in the next year! We’ll both need it!",
	}

	sliceClassicQuotes := []string{
		"“Always bear in mind that your own resolution to succeed is more important than any other.” -Abraham Lincoln",
		"“If you asked me for my New Year Resolution, it would be to find out who I am.” -Cyril Cusack",
		"“Let our New Year’s resolution be this: we will be there for one another as fellow members of humanity, in the finest sense of the word.” -Goran Persson",
		"“How few there are who have courage enough to own their faults, or resolution enough to mend them.” -Benjamin Franklin",
		"“Character is the ability to carry out a good resolution long after the excitement of the moment has passed.” -Cavett Robert",
		"“Each year’s regrets are envelopes in which messages of hope are found for the New Year.” -John R. Dallas Jr.",
		"“We must always change, renew, rejuvenate ourselves; otherwise, we harden.” -Johann Wolfgang von Goethe",
		"“One resolution I have made and try always to keep is this: To rise above the little things.” -John Burroughs",
		"“Good resolutions are simply checks that men draw on a bank where they have no account.” -Oscar Wilde",
		"“Learn from yesterday, live for today, hope for tomorrow.” -Albert Einstein",
		"“We spend January 1 walking through our lives, room by room, drawing up a list of work to be done, cracks to be patched. Maybe this year, to balance the list, we ought to walk through the rooms of our lives… not looking for flaws, but for potential.” -Ellen Goodman",
	}
	sliceFunnyQuotes := []string{
		"“Making resolutions is a cleansing ritual of self-assessment and repentance that demands personal honesty and, ultimately, reinforces humility. Breaking them is part of the cycle.” -Eric Zorn",
		"“A cat’s New Year dream is mostly a bird! Don’t be like a cat; in your New Year dream something that you have never dreamed. Target for new things.” -Mehmet Murat Ildan",
		"“My New Year’s Resolution List usually starts with the desire to lose between ten and three thousand pounds.” -Nia Vardalos",
		"“An optimist stays up till midnight to see the New Year in. A pessimist stays up to make sure the old year leaves.” -Bill Vaughn",
		"“May all your troubles last as long as your New Year’s resolutions.” -Joey Adams",
		"“A New Year’s resolution is something that goes in one Year and out the other.” -Anonymous",
		"“It wouldn’t be New Year’s if I didn’t have regrets.” -William Thomas",
		"“Never tell your resolution beforehand, or it’s twice as onerous a duty.” -John Selden",
		"“My New Year’s resolution is to stop hanging out with people who ask me about my New Year’s resolutions.” -Anonymous",
		"“I think I made too many New Year’s resolutions this year. It took me almost a full day to break them all.” -Anonymous",
		"“I can’t believe it’s been a year since I didn’t become a better person.” -Anonymous",
	}
	sliceMotivationalQuotes := []string{
		"“Write it on your heart that every day is the best day in the year.” -Ralph Waldo Emerson",
		"”Cheers to a new year and another chance for us to get it right.” -Oprah Winfrey",
		"“Tomorrow, is the first blank page of a 365-page book. Write a good one.” -Brad Paisley",
		"“One thing with gazing too frequently into the past is that we may turn around to find the future has run out on us.” -Michael Cibeuko",
		"“Year’s end is neither an end nor a beginning but a going on, with all the wisdom that experience can instill in us.” -Hal Borland",
		"“You are never too old to set another goal or to dream a new dream.” -C.S. Lewis",
		"“This is a new year. A new beginning. And things will change.” -Taylor Swift",
		"“Hope smiles from the threshold of the year to come, Whispering ‘it will be happier.” -Alfred Tennyson",
		"“For last year’s words belong to last year’s language and next year’s words await another voice.” -T.S. Eliot",
		"“We will open the book. Its pages are blank. We are going to put words on them ourselves. The book is called Opportunity, and its first chapter is New Year’s Day.” -Edith Lovejoy Pierce",
		"“Drop the last year into the silent limbo of the past. Let it go, for it was imperfect, and thank God that it can go.” -Brooks Atkinson",
		"“The object of a new year is not that we should have a new year. It is that we should have a new soul.” -G.K. Chesterton",
		"“Be always at war with your vices, at peace with your neighbors, and let each new year find you a better man.” -Benjamin Franklin",
		"“The bad news is time flies. The good news is you’re the pilot.” -Michael Altshuler",
		"“Although no one can go back and make a brand new start, anyone can start from now and make a brand new ending.” -Carl Bard",
	}
	sliceToast := []string{
		"“Here’s a toast to the future, a toast to the past, and a toast to our friends, far and near. The past a bright dream; may our friends remain faithful and clear.”",
		"“May you live as long as you want and never want as long as you live!”",
		"“As we start the New Year, let’s get down on our knees to thank God we’re on our feet.”",
		"“Here’s to a bright New Year and a fond farewell to the old; here’s to the things that are yet to come, and to the memories that we hold.”",
		"“May you have a prosperous New Year.”",
		"“Wishing you a happy, healthy New Year.”",
		"“May the New Year bless you with health, wealth, and happiness.”",
		"“In the New Year, may your right hand always be stretched out in friendship, never in want.”",
		"“May your teeth be white, your eyes be bright, and your capacity for love at its height!”",
		"“May the New Year see you loving, giving, and living!”",
		"“Happy New Year now and always!”",
		"“May your most used attire in the New Year be a smile.”",
		"“Have a wonderful New Year surrounded by those that mean the most!”",
		"“May you fill your New Year with new adventures, accomplishments, and learnings!”",
	}

	numberForInspirationalWish, err := getRandomNumber(len(sliceWishesInspirational) - 1)
	if err != nil {
		log.Fatal(err)
	}
	numberForClassicWish, err := getRandomNumber(len(sliceWishesClassic) - 1)
	if err != nil {
		log.Fatal(err)
	}
	numberForFunnyWish, err := getRandomNumber(len(sliceWishesFunny) - 1)
	if err != nil {
		log.Fatal(err)
	}
	numberForClassicQuotes, err := getRandomNumber(len(sliceClassicQuotes) - 1)
	if err != nil {
		log.Fatal(err)
	}
	numberForFunnyQuotes, err := getRandomNumber(len(sliceFunnyQuotes) - 1)
	if err != nil {
		log.Fatal(err)
	}
	numberForMotivationalQuotes, err := getRandomNumber(len(sliceMotivationalQuotes) - 1)
	if err != nil {
		log.Fatal(err)
	}
	numberForToast, err := getRandomNumber(len(sliceToast) - 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("---------------------------------------------")
	fmt.Println("")
	fmt.Printf("Inspirational Wish: %s! %v\n", customerFirstName, sliceWishesInspirational[numberForInspirationalWish])
	fmt.Println("")
	fmt.Printf("Funny Wish: %s! %v\n", customerFirstName, sliceWishesFunny[numberForFunnyWish])
	fmt.Println("")
	fmt.Printf("Classic Wish: %s! %v\n", customerFirstName, sliceWishesClassic[numberForClassicWish])
	fmt.Println("")
	fmt.Printf("%s! Classic quote: %v\n", customerFirstName, sliceClassicQuotes[numberForClassicQuotes])
	fmt.Println("")
	fmt.Printf("%s! Funny quote: %v\n", customerFirstName, sliceFunnyQuotes[numberForFunnyQuotes])
	fmt.Println("")
	fmt.Printf("%s! Motivational quote: %v\n", customerFirstName, sliceMotivationalQuotes[numberForMotivationalQuotes])
	fmt.Println("")
	fmt.Printf("%s! Your toast: %v\n", customerFirstName, sliceToast[numberForToast])
	fmt.Println("")
	fmt.Println("---------------------------------------------")
}
