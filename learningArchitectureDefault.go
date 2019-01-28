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
		Practices: []*models.Practice{
			&models.Practice{
				Name: "Agile",
			},
			&models.Practice{
				Name: "User-Centric Design",
			},
			&models.Practice{
				Name: "Open by Default",
			},
			&models.Practice{
				Name: "Open Standards & Solutions",
			},
			&models.Practice{
				Name: "Security & Privacy",
			},
			&models.Practice{
				Name: "Accessibility",
			},
			&models.Practice{
				Name: "Empowering People",
			},
			&models.Practice{
				Name: "Ethical & Responsible Use",
			},
			&models.Practice{
				Name: "Collaboration",
			},
		},
	},
	// Stream 2
	models.Stream{
		Name: "Design",
		Image: &models.Image{
			Path: "https://s3.amazonaws.com/foundationsapp/static/design.jpg",
		},
		Description: "User experience design (UX, UXD, UED or XD) is the process of enhancing user satisfaction with a product by improving the usability, accessibility, and pleasure provided in the interaction with the product.",
		Practices: []*models.Practice{
			&models.Practice{
				Name: "Design Thinking",
			},
			&models.Practice{
				Name: "Design Research",
			},
			&models.Practice{
				Name: "Content Design",
			},
			&models.Practice{
				Name: "Information Architecture",
			},
			&models.Practice{
				Name: "Service Design",
			},
			&models.Practice{
				Name: "User Interface",
			},
			&models.Practice{
				Name: "Interaction Design",
			},
			&models.Practice{
				Name: "Data Visualization",
			},
			&models.Practice{
				Name: "Usability Testing",
			},
			&models.Practice{
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
		Practices: []*models.Practice{
			&models.Practice{
				Name: "User-Centred Service Design",
			},
			&models.Practice{
				Name: "Leading Agile Teams",
			},
			&models.Practice{
				Name: "Leading Agile Projects",
			},
			&models.Practice{
				Name: "Leading Change",
			},
			&models.Practice{
				Name: "Communications",
			},
			&models.Practice{
				Name: "Disruptive Trends",
			},
			&models.Practice{
				Name: "Digital Governance",
			},
			&models.Practice{
				Name: "Agile Sponsorship",
			},
			&models.Practice{
				Name: "Fostering Innovation",
			},
			&models.Practice{
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
		Practices: []*models.Practice{
			&models.Practice{
				Name: "Biotechnology",
			},
			&models.Practice{
				Name: "Intelligence / Cognitive Augmentation",
			},
			&models.Practice{
				Name: "Foresight",
			},
			&models.Practice{
				Name: "IoT / Networks",
			},
			&models.Practice{
				Name: "3d Printing",
			},
			&models.Practice{
				Name: "Drones / Robotics",
			},
			&models.Practice{
				Name: "Blockchain / Distributed Systems",
			},
			&models.Practice{
				Name: "VR / Augmented Reality",
			},
			&models.Practice{
				Name: "AI",
			},
			&models.Practice{
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
		Practices: []*models.Practice{
			&models.Practice{
				Name: "Data Access",
			},
			&models.Practice{
				Name: "Data Cleaning",
			},
			&models.Practice{
				Name: "Data Manipulation",
			},
			&models.Practice{
				Name: "Natural Language Processing",
			},
			&models.Practice{
				Name: "Network Analysis",
			},
			&models.Practice{
				Name: "Geo-informatics",
			},
			&models.Practice{
				Name: "Statistical Analysis",
			},
			&models.Practice{
				Name: "Data Visualization",
			},
			&models.Practice{
				Name: "Streaming Data",
			},
			&models.Practice{
				Name: "Storytelling",
			},
		},
	},
	// Stream 6
	models.Stream{
		Name: "DevOps",
		Image: &models.Image{
			Path: "https://s3.amazonaws.com/foundationsapp/static/devops.jpg",
		},
		Description: "DevOps is the combination of cultural philosophies, practices, and tools that increases an organization's ability to deliver applications and services at high velocity: evolving and improving products at a faster pace than organizations using traditional software development and infrastructure management processes.",
		Practices: []*models.Practice{
			&models.Practice{
				Name: "Cloud Services",
			},
			&models.Practice{
				Name: "APIs",
			},
			&models.Practice{
				Name: "Automation",
			},
			&models.Practice{
				Name: "Testing",
			},
			&models.Practice{
				Name: "Containers",
			},
			&models.Practice{
				Name: "Cluster-Computing",
			},
			&models.Practice{
				Name: "Micro-services",
			},
		},
	},
	// Stream 7
}
