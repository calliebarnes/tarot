package main

import (
	"errors"
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
		// Major Arcana
		{"The Fool", "A joyful puppy, excited to embark on a thrilling journey.", "Welcome new adventures with a curious and open heart. \nTrust the path and embrace life's surprises.", "A foolish puppy, unaware of the dangers ahead.", "Be cautious, you might be heading towards a pitfall unknowingly.", false},
		{"The Magician", "A resourceful kitten, skillfully weaving its magic.", "Harness your inner abilities to manifest your dreams. \nUse your creativity to overcome obstacles.", "A mischievous kitten, playing tricks and causing mischief.", "Beware of manipulation, and remember to use your talents wisely.", false},
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

		// Minor Arcana
		// Suit of Cups
		{"Ace of Cups", "A jubilant dolphin, arcing gracefully above a shimmering sea, surrounded by sparkling droplets.", "Embrace budding emotions. Love, creativity, and fresh connections are blossoming.", "A dolphin trapped in a tight lagoon, longing for the open ocean.", "Explore your feelings deeply; potential emotional connections might be overlooked.", false},
		{"Two of Cups", "Two swans, their necks forming a heart against the backdrop of a radiant sunset.", "Nurture mutual bonds. A time of deepening relationships and heart-to-heart connections is at hand.", "Two swans, backs to each other, floating under a gloomy sky.", "Rekindle communication; differences need understanding and patience.", false},
		{"Three of Cups", "Three playful otters, linking paws and swirling in joyous dance upon shimmering waters.", "Relish shared moments. Celebrations, reunions, and collective triumphs await.", "Three otters, each facing away, lost in individual thoughts.", "Seek unity; cherish shared values and memories to maintain harmony.", false},
		{"Four of Cups", "A contemplative panda, seated beneath a tree, unaware of a radiant cup hovering nearby.", "Opportunities beckon. Ponder deeply, but remain receptive to external offers.", "A panda fixated on a tarnished cup, ignoring three golden ones.", "Broaden your horizon; variety and new experiences might hold unexpected joys.", false},
		{"Five of Cups", "A forlorn raccoon, gazing at toppled cups, the spilled liquid shimmering under moonlight.", "Reflect on past losses, but look ahead. Lessons learned pave the way for future gains.", "A raccoon overwhelmed by shadows, blind to the standing cups.", "Focus on healing and perspective; dwelling too much on past regrets hinders growth.", false},
		{"Six of Cups", "Joyful kittens, tumbling amidst vibrant flowers, reminiscent of carefree days.", "Rekindle childhood memories. Familiar faces and places offer comfort and insight.", "Kittens amidst withered blooms, lost in sepia-toned memories.", "Balance past joys with present realities. Over-nostalgia might cloud current joys.", false},
		{"Seven of Cups", "A mesmerized fox, gazing at a sky filled with vivid, dreamlike visions.", "Let imagination run wild, but discern dreams from reality. Infinite possibilities beckon.", "A fox, dizzied amidst a fog of illusions, struggling to focus.", "Stay grounded; every shimmering prospect might not align with your true path.", false},
		{"Eight of Cups", "A determined turtle, forging ahead, leaving familiar shores for mysterious horizons.", "Pursue deeper meaning. Some journeys require leaving the known behind.", "A turtle hesitating, gazing longingly back at past comforts.", "Seek progress; venturing beyond the familiar can unveil hidden treasures.", false},
		{"Nine of Cups", "A proud cat lounging amidst a bounty of overflowing, radiant cups.", "Savor the fruits of your labor. A time of contentment and personal fulfillment draws near.", "A cat, possessive of its hoard, wary of sharing.", "True joy often lies in sharing. Embrace generosity and avoid becoming too insular.", false},
		{"Ten of Cups", "Birds in harmonious flight, painting arcs of joy against a rainbow sky.", "Cherish shared happiness. Familial and communal bonds offer unparalleled joy.", "Birds silhouetted against gathering storm clouds, discord evident.", "Navigate misunderstandings; shared aspirations and mutual understanding will clear the air.", false},
		{"Page of Cups", "A curious seal, balancing a radiant cup on its nose, bubbles of wonder around.", "Embrace new emotional insights. Let your intuition and dreams guide you.", "A seal, puzzled, watching its cup teeter.", "Dive into unfamiliar feelings. Hidden depths offer profound understanding.", false},
		{"Knight of Cups", "A majestic heron, bearing a gleaming cup, soars towards dreamy horizons.", "Chase heartfelt desires. Passion and romance guide you, but ensure you remain anchored.", "A heron, drifting aimlessly, cup spilled.", "Direct your emotional energies with intent. Beware of becoming too whimsical or flighty.", false},
		{"Queen of Cups", "A graceful deer, standing by a serene pond, the moon's reflection in a radiant cup.", "Nurture emotional depths. Your empathy and intuition are powerful allies.", "A startled deer, its reflection distorted in rippled waters.", "Stay centered amidst emotional tides. Inner calm offers clarity amidst chaos.", false},
		{"King of Cups", "A wise bear, seated majestically, balancing a chalice amidst turbulent waters.", "Master your emotions. Lead with compassion, blending intellect with heart.", "A bear, perturbed, its cup tilting amidst stormy waves.", "Cultivate inner equilibrium. External storms require a steady heart and clear mind.", false},
		// Suit of Wands
		{"Ace of Wands", "A fiery dragon clutching a single wand, representing the spark of inspiration.", "New beginnings brim with potential. Embrace the creative spark within you.", "A dragon with its flame extinguished, holding a smoldering wand.", "Reignite your passion. The spark of inspiration is within reach, but overlooked.", false},
		{"Two of Wands", "A fox gazing thoughtfully over a vast landscape, two wands crossed behind.", "Plan your next steps with vision and foresight. The world is in your paws.", "A fox trapped between two towering wands, unable to move.", "Indecision may be your greatest barrier. Choose your path with care.", false},
		{"Three of Wands", "Three cheetahs sprinting forward, their path lined with wands, symbolizing rapid progress.", "Your efforts are coming to fruition. Prepare for swift and positive changes.", "Cheetahs at rest, ignoring the path ahead lined with wands.", "Delays are part of the journey. Patience and reflection are necessary.", false},
		{"Four of Wands", "Four zebras frolicking under a canopy formed by wands, a symbol of celebration.", "A time for joy and communal celebration. Share your successes with others.", "A lone zebra stands apart from a broken canopy of wands.", "Reconnect with your community. Unity and support may seem distant but are vital.", false},
		{"Five of Wands", "A group of squirrels in a playful tussle, wands scattered around, depicting friendly competition.", "Challenges ahead require your energy and wit. Approach them as opportunities for growth.", "Squirrels in disarray, their play turned to chaos amidst broken wands.", "Avoid unnecessary squabbles. Focus on constructive solutions.", false},
		{"Six of Wands", "A majestic stag with wands adorning its antlers, standing proudly on a hill.", "Recognition and success are on the horizon. Your leadership inspires others.", "A stag with its head bowed, wands fallen and scattered.", "Seek strength within. External validation may be fleeting.", false},
		{"Seven of Wands", "A mongoose stands defiantly atop a hill, bravely facing down encroaching wands.", "Defend your position with courage. Your perseverance will be your shield.", "A mongoose overwhelmed by shadows cast by towering wands.", "Reassess your battles. Knowing when to retreat can be a form of bravery.", false},
		{"Eight of Wands", "A flock of swallows darting swiftly through the air, wands parallel to their flight.", "Events are unfolding rapidly. Embrace the pace but keep your direction clear.", "Swallows caught in a tumultuous wind, their formation scattered.", "Communication may be muddled. Clarify your intentions and regroup.", false},
		{"Nine of Wands", "A weary but vigilant meerkat standing guard amidst a row of wands.", "Persistence is key. Stay alert and prepared, as the end goal is within reach.", "A meerkat slumped, its back turned to the wands.", "Consider seeking support. Your vigilance is commendable, but everyone needs rest.", false},
		{"Ten of Wands", "A camel laboriously carrying a bundle of wands across the desert, symbolizing burden.", "The weight of your responsibilities is heavy. Prioritize and seek help if needed.", "A camel collapsed under the weight of the wands, unable to proceed.", "Acknowledge your limits. It's time to lighten your load for your well-being.", false},
		{"Page of Wands", "A playful monkey swinging from a wand, curiosity shining in its eyes.", "New adventures beckon. Approach them with curiosity and a sense of play.", "A monkey sits at the base of a wand, its curiosity dimmed.", "Rekindle your sense of wonder. Opportunities for growth are all around.", false},
		{"Knight of Wands", "A spirited horse galloping fiercely, a wand clutched in its teeth, embodying momentum.", "Pursue your passions with full force. Your energy and determination will lead you.", "A horse hesitating, the wand dropped from its mouth.", "Reevaluate your direction. Passion without focus can lead to missed opportunities.", false},
		{"Queen of Wands", "A majestic lioness, her mane intertwined with wands, exuding calm authority.", "Lead with confidence and warmth. Your inner strength inspires and nurtures others.", "A lioness retreats into the shadows, her wands ignored.", "Reconnect with your inner fire. Leadership requires presence and engagement.", false},
		{"King of Wands", "An imposing eagle perched atop a wand, surveying the land with keen eyes.", "Your vision and decisiveness set the course. Lead with wisdom and courage.", "An eagle with its wings folded, perched low amidst fallen wands.", "Reflect on your leadership style. True strength lies in adaptability and wisdom.", false},
		// Suit of Swords
		{"Ace of Swords", "An eagle soaring high, clutching a gleaming sword in its talons, symbolizing clarity and breakthrough.", "Embrace the power of a new perspective. Truth and clarity cut through confusion.", "An eagle entangled, struggling to fly with a sword too heavy.", "Reevaluate your current challenges. The truth may be obscured or too complex to handle alone.", false},
		{"Two of Swords", "Two owls perched on crossed swords under a moonlit sky, representing a stalemate.", "A moment of decision looms. Balance your thoughts and intuition to choose wisely.", "Two owls turning their backs to each other, swords clashing below.", "Indecision may lead to conflict. Communication and compromise are needed.", false},
		{"Three of Swords", "A trio of swans entangled in thorny vines, swords piercing through, depicting heartache.", "Acknowledge the pain of loss or betrayal. Healing begins with acceptance.", "Swans separated by a storm, swords scattered in disarray.", "Avoid dwelling on sorrow. The storm will pass, and reconciliation is possible.", false},
		{"Four of Swords", "A sleeping fox in a peaceful clearing, swords laid down to signify rest.", "Take time to recover and find peace. Rest is necessary for clarity and progress.", "A restless fox, eyes open, surrounded by looming swords.", "Restlessness and worry may be disrupting your peace. Seek calm and detachment.", false},
		{"Five of Swords", "A cunning raven holding a sword in its beak, other swords abandoned below.", "Conflict may bring victory, but consider the cost. True triumph is in resolution.", "A raven flying away, leaving chaos and scattered swords behind.", "Avoid winning at all costs. The aftermath of conflict may bring isolation.", false},
		{"Six of Swords", "A turtle gently navigating through calm waters, swords laid flat in its path, symbolizing transition.", "Move forward with hope. Leaving troubles behind, smoother waters await.", "A turtle struggling against the tide, swords obstructing its path.", "Progress may be slow. Patience and perseverance will guide you through.", false},
		{"Seven of Swords", "A sly fox sneaking away with a sword, leaving others behind, indicating deceit.", "Be wary of deception, either from others or self-inflicted. Strategy is key.", "A fox caught in the act, its path blocked by the remaining swords.", "Dishonesty may backfire. Reassess your approach and strive for integrity.", false},
		{"Eight of Swords", "A butterfly ensnared in a web of swords and vines, symbolizing entrapment.", "Feeling trapped is often a state of mind. Seek new perspectives to find a way out.", "A butterfly breaking free, leaving behind the tangled swords.", "Liberation is within reach. Recognize your power to escape limiting beliefs.", false},
		{"Nine of Swords", "A nightingale singing amidst a storm, swords looming overhead, depicting worry.", "Anxiety may be overwhelming. Sharing your burdens can lighten your load.", "A silent nightingale, its song stifled by the oppressive weight of the swords.", "Silence won't ease your fears. Seek comfort in expression and connection.", false},
		{"Ten of Swords", "A fallen stag with swords embedded, a dark sky above, symbolizing betrayal.", "Endings may be painful, but they herald new beginnings. Release the past to heal.", "A stag rising, the swords falling away, dawn breaking.", "Recovery is possible. From the ashes of defeat, resilience is born.", false},
		{"Page of Swords", "A playful kitten batting at dangling swords, curious and eager to learn.", "Embrace your curiosity. Seek knowledge and truth with the innocence of youth.", "A kitten entangled in a curtain, swords forgotten, symbolizing distraction.", "Stay focused. Curiosity is valuable, but don't lose sight of your goals.", false},
		{"Knight of Swords", "A charging wolf, its sharp focus mirrored in the sword it guards.", "Pursue your objectives with determination and focus. Let nothing deter your quest.", "A wolf at a crossroads, hesitating, its sword dropped.", "Reconsider your direction. Rash actions may lead to unforeseen consequences.", false},
		{"Queen of Swords", "An elegant swan with a crown of swords, wisdom shining in her eyes.", "Lead with clarity and intellect. Your insight cuts through deception and falsehood.", "A swan with its wings clipped, surrounded by dull swords.", "Reclaim your power. Even in times of adversity, your wisdom remains.", false},
		{"King of Swords", "A regal lion, its mane interwoven with swords, commanding respect and authority.", "Rule with fairness and intellect. Your decisions are guided by a clear, moral compass.", "A lion's roar silenced, its swords scattered and broken.", "Reflect on your leadership. True strength lies in wisdom and justice, not force.", false},
		//Suit of Pentacles
		{"Ace of Pentacles", "A mighty elephant lifting a glowing pentacle with its trunk, symbolizing prosperity and new financial opportunities.", "Embrace the abundance coming your way. A new path to material success is unfolding.", "An elephant with a pentacle slipping from its grasp, suggesting a missed opportunity.", "Stay alert to new opportunities. Prosperity is within reach, but requires your attention.", false},
		{"Two of Pentacles", "A playful monkey juggling two pentacles amidst a busy environment, representing balance in financial decisions.", "Balance is key. Adapt to changes and maintain equilibrium in your financial endeavors.", "A monkey overwhelmed, unable to keep the pentacles in the air.", "Find your grounding. Overextension may lead to instability.", false},
		{"Three of Pentacles", "Beavers working together to build a dam, each holding a pentacle, symbolizing collaboration and craftsmanship.", "Teamwork leads to success. Your skills and efforts are being recognized and rewarded.", "Beavers at odds, their dam unfinished, pentacles neglected.", "Collaboration is faltering. Realign on common goals to ensure the success of your projects.", false},
		{"Four of Pentacles", "A dragon hoarding pentacles, tightly coiled around them, indicating possessiveness over wealth.", "Security is important, but beware of becoming too guarded or miserly with your resources.", "A dragon with pentacles slipping through its scales, signifying loss through over-caution.", "Loosen your grip. Sharing and investing may lead to greater rewards.", false},
		{"Five of Pentacles", "A lone wolf in a snowy landscape, distant from the warmth of a lighted window, depicting financial loss or insecurity.", "Hard times may be upon you, but support is closer than you think. Don't be afraid to seek help.", "A wolf finding solace in the warmth of a newfound shelter.", "Relief from hardship is possible. Open yourself to assistance from unexpected places.", false},
		{"Six of Pentacles", "A generous stag with a flourishing antler, each point holding a pentacle, symbolizing generosity and sharing of wealth.", "Prosperity allows for generosity. Share your wealth, and it will return to you manifold.", "A stag with barren antlers, unable to support the pentacles.", "Reevaluate your distribution of resources. Fairness and balance are needed.", false},
		{"Seven of Pentacles", "A patient panda sitting beneath a bamboo shoot, watching pentacles grow, representing investment and the fruits of labor.", "Patience is a virtue. Your efforts will bear fruit in time; nurture your investments.", "A panda amidst withered bamboo, the pentacles unripe.", "Reassess your investments. Not all efforts lead to success; know when to redirect your energy.", false},
		{"Eight of Pentacles", "A diligent ant meticulously carving a pentacle, epitomizing skill development and attention to detail.", "Hone your craft. Dedication to your work will lead to recognition and success.", "An ant overwhelmed, surrounded by unfinished pentacles.", "Avoid burnout. Focus and dedication are key, but so is balance.", false},
		{"Nine of Pentacles", "A peacock with feathers adorned with pentacles, standing in a lush garden, representing well-deserved success and self-sufficiency.", "Relish in the rewards of your hard work. Enjoy the luxury and independence you've earned.", "A peacock with dulled feathers, the garden untended.", "Self-care is needed. Success is not just financial but also personal well-being.", false},
		{"Ten of Pentacles", "A pack of wolves, each contributing to the strength of the pack, with pentacles in their den, symbolizing wealth and family legacy.", "You've built a lasting foundation. Your achievements will benefit not just you, but your lineage.", "A scattered pack, the den in disarray, pentacles neglected.", "Focus on what truly matters. Wealth is not just material but also the legacy and bonds we build.", false},
		{"Page of Pentacles", "A curious rabbit sniffing at a pentacle half-buried in the ground, representing an opportunity for prosperity yet to be unearthed.", "New opportunities are on the horizon. Stay curious and open to learning.", "A rabbit distracted by its surroundings, the pentacle unnoticed.", "Focus on your goals. Opportunities are present, but you must be attentive to seize them.", false},
		{"Knight of Pentacles", "A steadfast tortoise carrying a pentacle on its back, symbolizing reliability and a methodical approach to success.", "Progress may be slow, but your steady pace and thoroughness will ensure success.", "A tortoise stalled, its path blocked, the pentacle unattended.", "Reevaluate your path. Persistence is valuable, but adaptability is crucial.", false},
		{"Queen of Pentacles", "A nurturing doe in a verdant forest, pentacles hanging from the branches around her, embodying abundance and care.", "Your warmth and practicality create a haven of prosperity for those you care about.", "A doe in a sparse forest, the pentacles out of reach.", "Reconnect with your nurturing side. Abundance comes from caring and sharing, not just accumulating.", false},
		{"King of Pentacles", "A majestic buffalo standing firm on a field of pentacles, representing wealth, stability, and leadership.", "Your leadership brings stability and prosperity to those around you. Stand firm in your convictions.", "A buffalo wandering from an empty field, the pentacles buried and forgotten.", "Leadership requires presence. Stay engaged and mindful of your resources and responsibilities.", false},
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

func drawCards(deck *[]Card, n int) ([]Card, error) {
	if n > len(*deck) {
		return nil, errors.New(fmt.Sprintf("Cannot draw %d cards. Only %d cards left in the deck.", n, len(*deck)))
	}
	drawnCards := (*deck)[:n]
	*deck = (*deck)[:n]
	return drawnCards, nil
}

func main() {
	deck := createDeck()
	shuffleDeck(deck)
	numCardsToDraw := 6
	drawnCards, err := drawCards(&deck, numCardsToDraw)

	if err != nil {
		fmt.Println(err)
		return
	}

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
		} else {
			fmt.Println(card.description)
			fmt.Println()
			fmt.Println(card.meaning)
			fmt.Println()
		}
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - -")
	}
}
