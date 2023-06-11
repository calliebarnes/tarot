package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	name                string
	description         string
	meaning             string
	reversedDescription string
	reversedMeaning     string
	reversed            bool
}

func createDeck() []Card {
	deck := []Card{
		{"The Fool", "A joyful puppy, excited to embark on a thrilling journey.", "Welcome new adventures with a curious and open heart. \nTrust the path and embrace life's surprises.", "A foolish puppy, unaware of the dangers ahead.", "Be cautious, you might be heading towards a pitfall unknowingly.", false},
		{"The Magician", "A resourceful kitten, skillfully weaving its magic.", "Harness your inner abilities to manifest your dreams. \nUse your creativity to overcome obstacles.", "A mischievous kitten, playing tricks and causing mischief.",  "Beware of manipulation, and remember to use your talents wisely.", false},
		{"The High Priestess", "A perceptive owl, unveiling mysteries and hidden knowledge.", "Listen to your inner voice and trust your intuition. \nExplore the depths of your subconscious.", "A secretive owl, concealing its wisdom from the world.", "Don't ignore your intuition, strive to regain your personal harmony.", false},
		{"The Empress", "A nurturing cat, showering her kittens with love and warmth.", "Cultivate self-love and care for others. \nIndulge in the beauty of life.", "A neglectful cat, abandoning her kittens to fend for themselves.", "Take care not to neglect self-care or become overly dependent on others.", false},
		{"The Emperor", "A commanding dog, confidently ruling over its domain.", "Assert your authority and create order. \nStand firm in your decisions and lead with integrity.", "A tyrannical dog, abusing its power and authority.", "Avoid misuse of power and control. \nTry not to become an authoritarian.", false},
		{"The Hierophant", "A scholarly owl, imparting wisdom and guidance.", "Seek knowledge from mentors and traditions. \nEmbrace the teachings of the wise.", "A rebellious owl, rejecting the status quo and challenging authority.", "It might be a good time to question established rules or break conventions.", false},
		{"The Lovers", "Two smitten kittens, their love radiating like a beacon.", "Follow your heart and stay true to yourself. \nCherish deep connections and shared values.", "Two heartbroken kittens, their love torn apart by conflict.", "Ensure your values are aligned and balance is maintained in relationships.", false},
		{"The Chariot", "A brave lion, charging ahead with unwavering determination.", "Focus your energy and conquer your goals. \nDisplay courage and self-discipline in your pursuits.", "A reckless lion, charging ahead without caution.", "Be wary of losing control or becoming too aggressive. \nWatch out for possible obstacles in your path.", false},
		{"Strength", "A gentle lioness, her fierce heart protecting her cubs.", "Face challenges with inner strength and compassion. \nTrust your instincts and embrace your power.", "A timid lioness, her fierce heart cowering in fear.", "Remember to be secure in your abilities. \nDon't let your courage or inner strength wane.", false},
		{"The Hermit", "A contemplative owl, seeking solitude and inner wisdom.", "Turn inward and reflect on your life's journey. \nFind clarity through introspection and self-discovery.", "A lonely owl, shunning the world and its distractions.", "Isolation might seem comfortable, but don't withdraw too much from society.", false},
		{"Wheel of Fortune", "A playful kitten, spinning the wheel and trusting fate.", "Embrace life's cycles and trust the flow of destiny. \nAcknowledge the role of chance in your path.", "A mischievous kitten, spinning the wheel and causing chaos.", "Change is a part of life. \nBe prepared for it and remember that bad luck is only temporary.", false},
		{"Justice", "A discerning owl, weighing every decision with fairness.", "Act with integrity and seek the truth. \nStay objective and consider all perspectives.", "A biased owl, favoring one side over the other.", "Stand up for what's fair. \nDon't let injustice slide, and remember that accountability is crucial.", false},
		{"The Hanged Man", "A tranquil sloth, seeing the world from a unique angle.", "Pause and reevaluate your situation. \nEmbrace alternative viewpoints and be receptive to change.", "A restless sloth, unable to find peace and contentment.", "Stagnation is often a sign that change is needed. \nDon't be afraid to make necessary sacrifices.", false},
		{"Death", "An awe-inspiring phoenix, rising anew from its own ashes.", "Let go of the past and welcome transformation. \nEmbrace growth and seize fresh opportunities.", "A fearful phoenix, unable to overcome its own fears.", "Let go of what no longer serves you. \nResistance to change won't benefit you in the long run.", false},
		{"Temperance", "A harmonious owl, blending diverse elements with grace.", "Strive for balance and moderation in all aspects of life. \nCultivate patience and adaptability.", "A chaotic owl, unable to find harmony in its surroundings.", "Strive for balance in all things. \nDon't lose sight of your long-term vision.", false},
		{"The Devil", "A cunning raccoon, a reminder of life's temptations.", "Beware of distractions and unhealthy desires. \nPractice self-awareness and resist overindulgence.", "A playful raccoon, unable to resist its temptations.", "If you're feeling trapped or dealing with unhealthy relationships, it's time to take action. \nDon't let materialism or addiction control your life.", false},
		{"The Tower", "Resilient birds, rebuilding their nest after turbulent winds.", "Accept change and face unexpected challenges. \nAdapt to new circumstances with an open mind.", "A fearful bird, unable to face the storm.", "Change can be scary, but it's often for the best. \nDon't resist transformation.", false},
		{"The Star", "A luminous firefly, a beacon of hope in the darkness.", "Believe in yourself and reach for the stars. \nFind inspiration in your dreams and aspirations.", "A lonely firefly, unable to find its way in the darkness.", "Don't lose hope. \nEven in times of despair, remember to have faith and seek inspiration.", false},
		{"The Moon", "An enchanting wolf, howling at the moon, awakening your inner mystic.", "Trust your intuition and delve into the unknown. \nEmbrace your authentic self and seek hidden truths.", "A fearful wolf, howling at the moon, unable to face its inner demons.", "In times of confusion or fear, seek clarity and understanding.", false},
		{"The Sun", "A vibrant sunflower, basking in the golden light of life.", "Celebrate joy and warmth in all aspects of existence. \nWelcome new opportunities with a radiant spirit.", "A wilted sunflower, unable to find joy in life.", "Feeling down can be a sign of necessary change. \nRemember that success is within your reach.", false},
		{"Judgement", "A trumpeting elephant, heralding rebirth and renewal.", "Embrace change and be open to growth. \nEvaluate your choices and forge a new path with clarity.", "A fearful elephant, unable to face the unknown.", "Don't avoid self-examination. \nSelf-awareness is key to overcoming fear of change.", false},
		{"The World", "A thriving earthworm, connecting all life in a harmonious cycle.", "Acknowledge the interconnectedness of all beings. \nEmbrace unity and embark on a new chapter of life.", "A lonely earthworm, unable to find its place in the world.", "Strive for completion in all things. \nFeeling disjointed or failing to achieve your goals can be signs of needed change.", false},
	}
	return deck
}

func shuffleDeck(deck []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	for i := range deck {
		deck[i].reversed = rand.Intn(2) == 1
	}
}

func drawCards(deck []Card, n int) []Card {
	return deck[:n]
}

func main() {
	deck := createDeck()
	shuffleDeck(deck)
	numCardsToDraw := 3
	drawnCards := drawCards(deck, numCardsToDraw)

	for _, card := range drawnCards {
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - -")
		fmt.Println()
		fmt.Println(card.name)
		if card.reversed {
			fmt.Println("Reversed")
			fmt.Println(card.reversedDescription)
			fmt.Println()
			fmt.Println(card.reversedMeaning)
			fmt.Println()
			fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - -")
		} else {
			fmt.Println(card.description)
			fmt.Println()
			fmt.Println(card.meaning)
			fmt.Println()
			fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - -")
		}
	}
}
