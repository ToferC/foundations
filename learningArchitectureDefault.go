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
		Description: "The use of digital technologies, as an integrated part of governments’ modernisation strategies, to create public value. Relies on a digital government ecosystem comprised of government actors, non-governmental organisations, businesses, citizens’ associations and individuals which supports the production of and access to data, services and content through interactions with the government.",
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
			"Inclusive Design": &models.Practice{
				Name: "Inclusive Design",
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
		Description: "Digital literacy is the ability to use a range of technological tools for varied purposes and to understand the digital environment sufficiently to make well-informed decisions and understand the art of the possible.",
		Practices: map[string]*models.Practice{
			"Using Information & Data": &models.Practice{
				Name: "Using Information & Data",
			},
			"Using Tools & Technology": &models.Practice{
				Name: "Using tools & technology",
			},
			"Community Building": &models.Practice{
				Name: "Community Building",
			},
			"Critical Thinking & Evaluation": &models.Practice{
				Name: "Critical Thinking & Evaluation",
			},
			"Digital Citizenship": &models.Practice{
				Name: "Digital Citizenship",
			},
			"Creativity & Problem Solving": &models.Practice{
				Name: "Creativity & Problem Solving",
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
			"Inclusive Design": &models.Practice{
				Name: "Inclusive Design",
			},
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
			"Visual Design": &models.Practice{
				Name: "Visual Design",
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
			"Data Collection": &models.Practice{
				Name: "Data Collection",
			},
			"Data Cleaning": &models.Practice{
				Name: "Data Cleaning",
			},
			"Data Manipulation": &models.Practice{
				Name: "Data Manipulation",
			},
			"Data Modeling": &models.Practice{
				Name: "Data Manipulation",
			},
			"Programming": &models.Practice{
				Name: "Programming",
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
			"Cloud Services": &models.Practice{
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
			"Cluster-Computing": &models.Practice{
				Name: "Cluster-Computing",
			},
			"Micro-Services": &models.Practice{
				Name: "Micro-services",
			},
		},
	},
	// Stream
	models.Stream{
		Name: "Development",
		Image: &models.Image{
			Path: "https://s3.amazonaws.com/foundationsapp/static/coding.jpeg",
		},
		Description: "Software development is the process of conceiving, specifying, designing, programming, documenting, testing, and bug fixing involved in creating and maintaining applications, frameworks, or other software components.",
		Practices: map[string]*models.Practice{
			"Open Source": &models.Practice{
				Name: "Open Source",
			},
			"Git & GitHub": &models.Practice{
				Name: "Git & GitHub",
			},
			"Command Line": &models.Practice{
				Name: "Command Line",
			},
			"Environments": &models.Practice{
				Name: "Environments",
			},
			"Packages & Libraries": &models.Practice{
				Name: "Packages & Libraries",
			},
			"Programming Languages": &models.Practice{
				Name: "Programming Languages",
			},
			"Web Development": &models.Practice{
				Name: "Web Development",
			},
			"Mobile Development": &models.Practice{
				Name: "Mobile Development",
			},
			"Front-End Development": &models.Practice{
				Name: "Front-End Development",
			},
			"Back-End Development": &models.Practice{
				Name: "Back-End Development",
			},
			"Databases": &models.Practice{
				Name: "Databases",
			},
			"APIs": &models.Practice{
				Name: "APIs",
			},
		},
	},
}

var skillMap = map[int]string{
	1: "Beginner",
	2: "Experienced",
	3: "Professional",
	4: "Expert",
}
