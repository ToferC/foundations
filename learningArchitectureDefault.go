package main

import (
	"github.com/toferc/foundations/models"
)

var baseArchitecture = []models.Stream{
	models.Stream{
		// Stream 1
		Name: "Digital Government",
		Image: &models.Image{
			Path: "https://s3.amazonaws.com/foundationsapp/static/digital+government.png",
		},
		Description: "“The use of digital technologies, as an integrated part of governments’ modernisation strategies, to create public value. Relies on a digital government ecosystem comprised of government actors, non-governmental organisations, businesses, citizens’ associations and individuals which supports the production of and access to data, services and content through interactions with the government.",
		Practices: map[string]*models.Practice{
			"Agile": &models.Practice{
				Name: "Agile",
			},
			"User-Centric Design": &models.Practice{
				Name: "User-Centric Design",
			},
			"Open by Default": &models.Practice{
				Name: "Open by Default",
			},
			"Open Standards & Solutions": &models.Practice{
				Name: "Open Standards & Solutions",
			},
			"Security & Privacy": &models.Practice{
				Name: "Security & Privacy",
			},
			"Accessibility": &models.Practice{
				Name: "Accessibility",
			},
			"Empowering People": &models.Practice{
				Name: "Empowering People",
			},
			"Ethical & Responsible Use": &models.Practice{
				Name: "Ethical & Responsible Use",
			},
			"Collaboration": &models.Practice{
				Name: "Collaboration",
			},
		},
	},
	// Stream
	models.Stream{
		Name: "Digital Literacy",
		Image: &models.Image{
			Path: "https://s3.amazonaws.com/foundationsapp/static/digital_literacy.jpg",
		},
		Description: "the ability to use a range of technological tools for varied purposes.",
		Practices: map[string]*models.Practice{
			"Creativity & Collaboration": &models.Practice{
				Name: "Creativity & Collaboration",
			},
			"Curation": &models.Practice{
				Name: "Curation",
			},
			"Understanding Technology": &models.Practice{
				Name: "Understanding Technology",
			},
			"Critical Thinking": &models.Practice{
				Name: "Critical Thinking",
			},
			"Cultural & Social Impacts": &models.Practice{
				Name: "Cultural & Social Impacts",
			},
			"Problem Solving": &models.Practice{
				Name: "Problem Solving",
			},
		},
	},
	// Stream
	models.Stream{
		Name: "Design",
		Image: &models.Image{
			Path: "https://s3.amazonaws.com/foundationsapp/static/design.jpg",
		},
		Description: "User experience design (UX, UXD, UED or XD) is the process of enhancing user satisfaction with a product by improving the usability, accessibility, and pleasure provided in the interaction with the product.",
		Practices: map[string]*models.Practice{
			"Design Thinking": &models.Practice{
				Name: "Design Thinking",
			},
			"Design Research": &models.Practice{
				Name: "Design Research",
			},
			"Content Design": &models.Practice{
				Name: "Content Design",
			},
			"Information Architecture": &models.Practice{
				Name: "Information Architecture",
			},
			"Service Design": &models.Practice{
				Name: "Service Design",
			},
			"User Interface": &models.Practice{
				Name: "User Interface",
			},
			"Interaction Design": &models.Practice{
				Name: "Interaction Design",
			},
			"Data Visualization": &models.Practice{
				Name: "Data Visualization",
			},
			"Usability Testing": &models.Practice{
				Name: "Usability Testing",
			},
			"Prototyping & Iteration": &models.Practice{
				Name: "Prototyping & Iteration",
			},
		},
	},
	// Stream 3
	models.Stream{
		Name: "Leadership",
		Image: &models.Image{
			Path: "https://s3.amazonaws.com/foundationsapp/static/leadership.jpg",
		},
		Description: "Digital leadership is the strategic use of a company's digital assets to achieve business goals. Digital leadership can be addressed at both organizational and individual levels.",
		Practices: map[string]*models.Practice{
			"User-Centred Service Design": &models.Practice{
				Name: "User-Centred Service Design",
			},
			"Leading Agile Teams": &models.Practice{
				Name: "Leading Agile Teams",
			},
			"Leading Agile Projects": &models.Practice{
				Name: "Leading Agile Projects",
			},
			"Leading Change": &models.Practice{
				Name: "Leading Change",
			},
			"Communications": &models.Practice{
				Name: "Communications",
			},
			"Disruptive Trends": &models.Practice{
				Name: "Disruptive Trends",
			},
			"Digital Governance": &models.Practice{
				Name: "Digital Governance",
			},
			"Agile Sponsorship": &models.Practice{
				Name: "Agile Sponsorship",
			},
			"Fostering Innovation": &models.Practice{
				Name: "Fostering Innovation",
			},
			"Openness & Collaboration": &models.Practice{
				Name: "Openness & Collaboration",
			},
		},
	},
	// Stream 4
	models.Stream{
		Name: "Disruptive Technology",
		Image: &models.Image{
			Path: "https://s3.amazonaws.com/foundationsapp/static/disruptive.png",
		},
		Description: " Disruptive technologies are those that significantly alter the way businesses or entire industries operate. Often times, these technologies force companies to alter the way they approach their business, or risk losing market share or becoming irrelevant.",
		Practices: map[string]*models.Practice{
			"Biotechnology": &models.Practice{
				Name: "Biotechnology",
			},
			"Intelligence / Cognitive Augmentation": &models.Practice{
				Name: "Intelligence / Cognitive Augmentation",
			},
			"Foresight": &models.Practice{
				Name: "Foresight",
			},
			"IoT / Networks": &models.Practice{
				Name: "IoT / Networks",
			},
			"3d Printing": &models.Practice{
				Name: "3d Printing",
			},
			"Drones / Robotics": &models.Practice{
				Name: "Drones / Robotics",
			},
			"Blockchain / Distributed Systems": &models.Practice{
				Name: "Blockchain / Distributed Systems",
			},
			"VR / Augmented Reality": &models.Practice{
				Name: "VR / Augmented Reality",
			},
			"AI": &models.Practice{
				Name: "AI",
			},
			"Micro / Nano-materials": &models.Practice{
				Name: "Micro / Nano-materials",
			},
		},
	},
	// Stream 5
	models.Stream{
		Name: "Data Analysis",
		Image: &models.Image{
			Path: "https://s3.amazonaws.com/foundationsapp/static/data.jpg",
		},
		Description: "Data analysis is a process of inspecting, cleansing, transforming, and modeling data with the goal of discovering useful information, informing conclusions, and supporting decision-making.",
		Practices: map[string]*models.Practice{
			"Data Access": &models.Practice{
				Name: "Data Access",
			},
			"Data Cleaning": &models.Practice{
				Name: "Data Cleaning",
			},
			"Data Manipulation": &models.Practice{
				Name: "Data Manipulation",
			},
			"Coding Basics": &models.Practice{
				Name: "Coding Basics",
			},
			"Pandas / DataFrames": &models.Practice{
				Name: "Pandas / DataFrames",
			},
			"Natural Language Processing": &models.Practice{
				Name: "Natural Language Processing",
			},
			"Network Analysis": &models.Practice{
				Name: "Network Analysis",
			},
			"Geo-informatics": &models.Practice{
				Name: "Geo-informatics",
			},
			"Statistical Analysis": &models.Practice{
				Name: "Statistical Analysis",
			},
			"Data Visualization": &models.Practice{
				Name: "Data Visualization",
			},
			"Streaming Data": &models.Practice{
				Name: "Streaming Data",
			},
			"Storytelling": &models.Practice{
				Name: "Storytelling",
			},
		},
	},
	// Stream
	models.Stream{
		Name: "AI / Machine Learning",
		Image: &models.Image{
			Path: "https://s3.amazonaws.com/foundationsapp/static/ai.jpg",
		},
		Description: "Machine learning is an application of artificial intelligence (AI) that provides systems the ability to automatically learn and improve from experience without being explicitly programmed. Machine learning focuses on the development of computer programs that can access data and use it learn for themselves.",
		Practices: map[string]*models.Practice{
			"Unsupervised Learning": &models.Practice{
				Name: "Unsupervised Learning",
			},
			"Supervised Learning": &models.Practice{
				Name: "Supervised Learning",
			},
			"Reinforcement Learning": &models.Practice{
				Name: "Reinforcement Learning",
			},
			"LSTMs": &models.Practice{
				Name: "LSTMs",
			},
			"GANs": &models.Practice{
				Name: "GANs",
			},
			"Deep Learning / Neural Networks": &models.Practice{
				Name: "Deep Learning / Neural Networks",
			},
			"Statistics": &models.Practice{
				Name: "Statistics",
			},
			"Streaming Data": &models.Practice{
				Name: "Streaming Data",
			},
			"Big Data": &models.Practice{
				Name: "Big Data",
			},
			"Bias & Ethics": &models.Practice{
				Name: "Bias & Ethics",
			},
		},
	},
	// Stream
	models.Stream{
		Name: "DevOps",
		Image: &models.Image{
			Path: "https://s3.amazonaws.com/foundationsapp/static/devops.jpg",
		},
		Description: "DevOps is the combination of cultural philosophies, practices, and tools that increases an organization's ability to deliver applications and services at high velocity: evolving and improving products at a faster pace than organizations using traditional software development and infrastructure management processes.",
		Practices: map[string]*models.Practice{
			"Cloud": &models.Practice{
				Name: "Cloud Services",
			},
			"APIs": &models.Practice{
				Name: "APIs",
			},
			"Automation": &models.Practice{
				Name: "Automation",
			},
			"Testing": &models.Practice{
				Name: "Testing",
			},
			"Containers": &models.Practice{
				Name: "Containers",
			},
			"Cluster": &models.Practice{
				Name: "Cluster-Computing",
			},
			"Micro": &models.Practice{
				Name: "Micro-services",
			},
		},
	},
}
