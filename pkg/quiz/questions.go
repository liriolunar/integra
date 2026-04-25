package quiz

type Question struct {
	Topic       string
	Prompt      string
	Options     []string
	Answer      int
	Explanation string
}

var Bank = []Question{
	{
		Topic:  "Limites",
		Prompt: "Qual é o valor de lim x→2 (x + 3) ?",
		Options: []string{
			"2",
			"3",
			"5",
			"6",
		},
		Answer:      2,
		Explanation: "Em funções contínuas, basta substituir x por 2. Assim, 2 + 3 = 5.",
	},
	{
		Topic:  "Limites",
		Prompt: "Qual é o limite lim x→0 de x² ?",
		Options: []string{
			"0",
			"1",
			"2",
			"não existe",
		},
		Answer:      0,
		Explanation: "Quando x tende a 0, x² também tende a 0.",
	},
	{
		Topic:  "Limites",
		Prompt: "Qual é o limite lim x→1 (2x + 1) ?",
		Options: []string{
			"1",
			"2",
			"3",
			"4",
		},
		Answer:      2,
		Explanation: "Como a função é contínua, basta substituir x por 1: 2(1) + 1 = 3.",
	},
	{
		Topic:  "Derivadas",
		Prompt: "Qual é a derivada de f(x) = x² ?",
		Options: []string{
			"2x",
			"x",
			"x³",
			"2",
		},
		Answer:      0,
		Explanation: "Pela regra da potência, d/dx (x²) = 2x.",
	},
	{
		Topic:  "Derivadas",
		Prompt: "Qual é a derivada de f(x) = 3x ?",
		Options: []string{
			"3",
			"x",
			"3x²",
			"0",
		},
		Answer:      0,
		Explanation: "A derivada de ax é a, então a derivada de 3x é 3.",
	},
	{
		Topic:  "Derivadas",
		Prompt: "Qual é a derivada de f(x) = x³ ?",
		Options: []string{
			"3x²",
			"x²",
			"3x",
			"x⁴",
		},
		Answer:      0,
		Explanation: "Pela regra da potência, d/dx (x³) = 3x².",
	},
	{
		Topic:  "Integrais",
		Prompt: "Pegadinha final: qual é o valor de ∫ e^(2x) dx ?",
		Options: []string{
			"e^(2x) + C",
			"2e^(2x) + C",
			"(1/2)e^(2x) + C",
			"e^x + C",
		},
		Answer:      2,
		Explanation: "A derivada de e^(2x) é 2e^(2x). Por isso, a primitiva correta é (1/2)e^(2x) + C.",
	},
}
