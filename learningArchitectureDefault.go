package main

import (
	"github.com/toferc/foundations/models"
)

var baseArchitecture = []models.Stream{
	models.Stream{
		// Stream 1
		Name: "Digital Government",
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
