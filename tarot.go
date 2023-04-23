package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	name        string
	description string
	meaning     string
}

func createDeck() []Card {
	deck := []Card{
		{"The Fool", "A joyful puppy, excited to embark on a thrilling journey.", "Welcome new adventures with a curious and open heart. \nTrust the path and embrace life's surprises."},
		{"The Magician", "A resourceful kitten, skillfully weaving its magic.", "Harness your inner abilities to manifest your dreams. \nUse your creativity to overcome obstacles."},
		{"The High Priestess", "A perceptive owl, unveiling mysteries and hidden knowledge.", "Listen to your inner voice and trust your intuition. \nExplore the depths of your subconscious."},
		{"The Empress", "A nurturing cat, showering her kittens with love and warmth.", "Cultivate self-love and care for others. \nIndulge in the beauty of life."},
		{"The Emperor", "A commanding dog, confidently ruling over its domain.", "Assert your authority and create order. \nStand firm in your decisions and lead with integrity."},
		{"The Hierophant", "A scholarly owl, imparting wisdom and guidance.", "Seek knowledge from mentors and traditions. \nEmbrace the teachings of the wise."},
		{"The Lovers", "Two smitten kittens, their love radiating like a beacon.", "Follow your heart and stay true to yourself. \nCherish deep connections and shared values."},
		{"The Chariot", "A brave lion, charging ahead with unwavering determination.", "Focus your energy and conquer your goals. \nDisplay courage and self-discipline in your pursuits."},
		{"Strength", "A gentle lioness, her fierce heart protecting her cubs.", "Face challenges with inner strength and compassion. \nTrust your instincts and embrace your power."},
		{"The Hermit", "A contemplative owl, seeking solitude and inner wisdom.", "Turn inward and reflect on your life's journey. \nFind clarity through introspection and self-discovery."},
		{"Wheel of Fortune", "A playful kitten, spinning the wheel and trusting fate.", "Embrace life's cycles and trust the flow of destiny. \nAcknowledge the role of chance in your path."},
		{"Justice", "A discerning owl, weighing every decision with fairness.", "Act with integrity and seek the truth. \nStay objective and consider all perspectives."},
		{"The Hanged Man", "A tranquil sloth, seeing the world from a unique angle.", "Pause and reevaluate your situation. \nEmbrace alternative viewpoints and be receptive to change."},
		{"Death", "An awe-inspiring phoenix, rising anew from its own ashes.", "Let go of the past and welcome transformation. \nEmbrace growth and seize fresh opportunities."},
		{"Temperance", "A harmonious owl, blending diverse elements with grace.", "Strive for balance and moderation in all aspects of life. \nCultivate patience and adaptability."},
		{"The Devil", "A cunning raccoon, a reminder of life's temptations.", "Beware of distractions and unhealthy desires. \nPractice self-awareness and resist overindulgence."},
		{"The Tower", "Resilient birds, rebuilding their nest after turbulent winds.", "Accept change and face unexpected challenges. \nAdapt to new circumstances with an open mind."},
		{"The Star", "A luminous firefly, a beacon of hope in the darkness.", "Believe in yourself and reach for the stars. \nFind inspiration in your dreams and aspirations."},
		{"The Moon", "An enchanting wolf, howling at the moon, awakening your inner mystic.", "Trust your intuition and delve into the unknown. \nEmbrace your authentic self and seek hidden truths."},
		{"The Sun", "A vibrant sunflower, basking in the golden light of life.", "Celebrate joy and warmth in all aspects of existence. \nWelcome new opportunities with a radiant spirit."},
		{"Judgement", "A trumpeting elephant, heralding rebirth and renewal.", "Embrace change and be open to growth. \nEvaluate your choices and forge a new path with clarity."},
		{"The World", "A thriving earthworm, connecting all life in a harmonious cycle.", "Acknowledge the interconnectedness of all beings. \nEmbrace unity and embark on a new chapter of life."},
	}
	return deck
}

func shuffleDeck(deck []Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

func drawCards(deck []Card, n int) []Card {
	return deck[:n]
}

func main() {
	deck := createDeck()
	shuffleDeck(deck)
	numCardsToDraw := 1
	drawnCards := drawCards(deck, numCardsToDraw)

	for _, card := range drawnCards {
		fmt.Println(card.name)
		fmt.Println(card.description)
		fmt.Println()
		fmt.Println(card.meaning)
	}
}
