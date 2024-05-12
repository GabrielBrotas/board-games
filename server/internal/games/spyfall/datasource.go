package spyfall

import (
	"math/rand"
)

type iLocationAndRoles struct {
	Location string
	slug     string
	Roles    []string
}

func getLocationAndRoles() []iLocationAndRoles {
	return []iLocationAndRoles{
		{
			Location: "Hospital",
			Roles:    []string{"Médico", "Enfermeiro", "Cirurgião", "Recepcionista", "Administrador", "Técnico de Laboratório", "Radiologista", "Psicólogo", "Nutricionista", "Pediatra", "Dermatologista", "Cardiologista", "Ortopedista", "Visitante", "Paciente"},
			slug:     "hospital",
		},
		{
			Location: "Estação Espacial",
			Roles:    []string{"Astronauta", "Engenheiro de Voo", "Comandante", "Físico", "Médico", "Especialista em Comunicações", "Cientista de Dados", "Psicólogo", "Meteorologista", "Engenheiro Robótico", "Químico", "Astrônomo", "Turista", "Piloto", "Visitante"},
			slug:     "estacao-espacial",
		},
		{
			Location: "Supermercado",
			Roles:    []string{"Caixa", "Gerente", "Repositor", "Açougueiro", "Padeiro", "Atendente de Frios", "Estoquista", "Empacotador", "Fiscal de Loja", "Limpeza", "Administrador", "Analista de Estoque", "Balconista", "Confeiteiro", "Jovem Aprendiz"},
			slug:     "supermercado",
		},
		{
			Location: "Submarino",
			Roles:    []string{"Capitão", "Engenheiro Chefe", "Operador de Sonar", "Mergulhador", "Técnico em Eletrônica", "Operador de Comunicações", "Cientista Marinho", "Analista de Sistemas", "Operador de Torpedo", "Especialista em Navegação", "Meteorologista", "Operador de Radar", "Visitante", "Diretor de Engenharia Submarina", "Gerente de Projetos"},
			slug:     "submarino",
		},
		{
			Location: "Banco",
			Roles:    []string{"Gerente", "Segurança", "Analista Financeiro", "Atendente", "Economista", "Ladrão", "Advogado", "Consultor de Investimentos", "Analista de Crédito", "Operador de Câmbio", "Especialista em Riscos", "Administrador de Patrimônio", "Limpeza", "Gerente de Contas", "Corretor de Seguros"},
			slug:     "banco",
		},
		{
			Location: "Escola",
			Roles:    []string{"Professor", "Diretor", "Coordenador Pedagógico", "Secretário", "Psicólogo Escolar", "Técnico de Informática", "Porteiro", "Limpeza", "Monitor de Educação Física", "Cozinheira", "Conselheiro Tutelar", "Pai de Aluno", "Motorista de Ônibus Escolar", "Professor de Ingles", "Professor de Matematica"},
			slug:     "escola",
		},
		{
			Location: "Circo",
			Roles:    []string{"Palhaço", "Malabarista", "Domador de Leões", "Acrobata", "Mágico", "Vendedor de Ingressos", "Trapezista", "Equilibrista", "Costureira de Figurinos", "Bilheteiro", "Fotógrafo", "Diretor Artístico", "Coreógrafo", "Maquiador", "Limpeza"},
			slug:     "circo",
		},
		{
			Location: "Restaurante",
			Roles:    []string{"Garçom", "Cozinheiro", "Barman", "Recepcionista", "Gerente", "Lavador de Pratos", "Churrasqueiro", "Pizzaiolo", "Limpeza", "Caixa", "Estoquista", "Chef de Cozinha", "Entregador", "Auxiliar de Cozinha", "Auxiliar de Limpeza"},
			slug:     "restaurante",
		},
		{
			Location: "Teatro",
			Roles:    []string{"Ator", "Diretor", "Cenógrafo", "Iluminador", "Maquiador", "Figurinista", "Bilheteiro", "Coreógrafo", "Técnico de Som e Luz", "Produtor", "Assistente de Palco", "Limpeza", "Fotógrafo", "Pintor de Cenários", "Cameraman"},
			slug:     "teatro",
		},
		{
			Location: "Aeroporto",
			Roles:    []string{"Piloto", "Comissário de Bordo", "Agente de Check-in", "Segurança", "Controlador de Tráfego Aéreo", "Mecânico de Aviação", "Carregador de Bagagem", "Limpeza", "Gerente de Operações", "Chefe de Cabine", "Recepcionista de Sala VIP", "Guia de Turismo", "Agente de Imigração", "Administrador de Aeroporto", "Profissional de Logistica"},
			slug:     "aeroporto",
		},
		{
			Location: "Zoológico",
			Roles:    []string{"Biólogo", "Veterinário", "Guia Turístico", "Cuidador de Animais", "Segurança", "Fotógrafo", "Administrador", "Atendente de Loja de Souvenirs", "Jardineiro", "Educador Ambiental", "Nutricionista de Animais", "Serviços Gerais", "Pesquisador", "Transportador de Animais", "Gestor Ambiental"},
			slug:     "zoo",
		},
		{
			Location: "Cassino",
			Roles:    []string{"Croupier", "Segurança", "Gerente", "Bartender", "Caixa", "Dealer de Poker", "Apostador", "Limpeza", "Administrador Financeiro", "Atendente de Estacionamento", "Recepcionista", "Contador", "Cliente", "Garçom", "Apostador"},
			slug:     "cassino",
		},
		{
			Location: "Navio Cruzeiro",
			Roles:    []string{"Capitão", "Marinheiro", "Cozinheiro", "Animador de Bordo", "Médico", "Barman", "Recepcionista", "Segurança", "Garçom", "Fotógrafo", "Guia de Excursões", "Dj", "Operador de Cabine", "Camareiro", "Salva-vidas"},
			slug:     "navio-cruzeiro",
		},
		{
			Location: "Parque de Diversões",
			Roles:    []string{"Operador de Brinquedo", "Bilheteiro", "Vendedor de Alimentos", "Palhaço", "Pintor de Rosto", "Atendente de Loja de Brinquedos", "Gerente de Parque", "Operador de Montanha-Russa", "Escultor de Balões", "Artista de Circo", "Técnico de Efeitos Especiais", "Especialista em Segurança de Brinquedos", "Artista de Maquiagem Teatral", "Designer de Parques Temáticos", "Operador de Carrosel"},
			slug:     "parque-diversoes",
		},
		{
			Location: "Museu",
			Roles:    []string{"Guia Turístico", "Segurança", "Designer de Exposições", "Pesquisador", "Conservador de Arte", "Restaurador de Documentos", "Arqueólogo", "Especialista em Digitalização", "Especialista em Conservação de Materiais", "Especialista em Restauração de Esculturas", "Educador de Museu", "Diretor de Museu", "Guia de Museu", "Visitante", "Gerente de Loja de Presentes"},
			slug:     "museu",
		},
		{
			Location: "Estúdio de TV",
			Roles:    []string{"Apresentador", "Figurinista", "Maquiador", "Engenheiro de Som", "Editor de Vídeo", "Ator", "Roteirista", "Assistente de Produção", "Operador de Câmera", "Jornalista", "Diretor de TV", "Assistente de Direção", "Coordenador de Elenco", "Técnico de Transmissão", "Gerente de Programação"},
			slug:     "estudio-tv",
		},
		{
			Location: "Avião",
			Roles:    []string{"Piloto", "Copiloto", "Comissário de Bordo", "Mecânico de Aviação", "Terrorista", "Coordenador de Voo", "Técnico de Segurança de Voo", "Instrutor de Voo", "Atendente de Serviços ao Passageiro", "Chefe de Cabine", "Agente de Limpeza de Aeronaves", "Passageiro Primeira Classe", "Passageiro Classe Economica", "Aeromoça", "Turista"},
			slug:     "aviao",
		},
		{
			Location: "Praia",
			Roles:    []string{"Salva-vidas", "Instrutor de Surf", "Vendedor Ambulante", "Gerente de Quiosque", "Fotógrafo de Praia", "Instrutor de Mergulho", "Operador de Passeio de Barco", "Banhista", "Instrutor de Vôlei de Praia", "Operador de Aluguel de Equipamentos", "Jogador de Altinha", "Surfista", "Ladrão", "Guia Turístico Local", "Visitante"},
			slug:     "praia",
		},
		{
			Location: "Cinema",
			Roles:    []string{"Bilheteiro", "Operador de Projeção", "Cliente", "Gerente de Cinema", "Técnico de Som", "Agente de Limpeza", "Segurança", "Programador de Filmes", "Assistente de Marketing", "Host de Estreias", "Designer Gráfico", "Analista de Atendimento ao Cliente", "Operador de Cabine", "Operador de Máquina de Pipoca", "Operador de Máquina de Refrigerante"},
			slug:     "cinema",
		},
		{
			Location: "Base Militar",
			Roles:    []string{"Comandante da Base", "Oficial de Operações", "Médico Militar", "Engenheiro Militar", "Sargento de Armas", "Controlador de Tráfego Aéreo Militar", "Técnico em Comunicações", "Instrutor de Combate", "Mecânico de Veículos Militares", "Soldado", "Cozinheiro Militar", "Sargento", "Recruta", "Coronel", "Sniper"},
			slug:     "base-militar",
		},
		{
			Location: "Spa",
			Roles:    []string{"Esteticista", "Gerente de Spa", "Instrutor de Yoga", "Especialista em Aromaterapia", "Terapeuta", "Instrutor de Pilates", "Gerente de Produtos de Spa", "Especialista em Terapias com Água", "Coordenador de Eventos de Bem-Estar", "Designer de Interiores de Spa", "Especialista em Banhos Terapêuticos", "Terapeuta Capilar", "Cliente", "Especialista em Terapias de Pedras Quentes", "Especialista em Terapias de Lama"},
			slug:     "spa",
		},
		{
			Location: "Trem",
			Roles:    []string{"Maquinista", "Condutor", "Fiscal de Trem", "Atendente de Bordo", "Técnico de Manutenção de Trilhos", "Agente de Bilheteria", "Engenheiro Ferroviário", "Chefe de Serviço de Passageiros", "Operador de Sala de Controle", "Segurança Ferroviária", "Coordenador de Logística", "Agente de Limpeza", "Atendente de Bar no Trem", "Assistente de Atendimento ao Cliente", "Passageiro"},
			slug:     "trem",
		},
		{
			Location: "Delegacia",
			Roles:    []string{"Delegado", "Investigador", "Agente de Polícia", "Operador de Câmeras de Segurança", "Administrador de Sistema Penitenciário", "Psicólogo Policial", "Especialista em Relações Comunitárias", "Instrutor de Armas de Fogo", "Motorista Policial", "Guarda de Cela", "Negociador de Reféns", "Coordenador de Programas de Prevenção ao Crime", "Detento", "Ladrão", "Testemunha"},
			slug:     "delegacia",
		},
		{
			Location: "Oficina",
			Roles:    []string{"Mecânico de Automóveis", "Eletricista Automotivo", "Pintor de Carros", "Técnico de Ar Condicionado Automotivo", "Gerente de Oficina", "Recepcionista de Oficina", "Estoquista de Peças", "Vendedor de Peças Automotivas", "Lavador de Carros", "Técnico em Alinhamento e Balanceamento", "Especialista em Restauração de Veículos Antigos", "Técnico em Eletrônica Automotiva", "Inspetor de Qualidade", "Consultor Técnico Automotivo", "Especialista em Pneus"},
			slug:     "oficina",
		},
		{
			Location: "Estádio de Futebol",
			Roles:    []string{"Gerente de Estádio", "Segurança de Estádio", "Coordenador de Eventos Esportivos", "Relações Públicas do Clube", "Técnico de Futebol", "Jogador de Futebol", "Fisioterapeuta Esportivo", "Narrador Esportivo", "Médico Esportivo", "Operador de Câmera", "Vendedor de Ingressos", "Diretor de Operações de Jogo", "Supervisor de Limpeza", "Organizador de Torcida", "Agente de Atendimento ao Cliente"},
			slug:     "estadio-futebol",
		},
	}
}

func generateLocationAndRoles() iLocationAndRoles {
	locationsAndRoles := getLocationAndRoles()

	// randomly select a location and shuffle the roles
	idx := rand.Intn(len(locationsAndRoles))
	selected := locationsAndRoles[idx]

	shuffleRoles(selected.Roles)

	return selected
}

func getRandomRole(roles []string) string {
	idx := rand.Intn(len(roles))
	return roles[idx]
}

func shuffleRoles(roles []string) {
	for i := range roles {
		j := rand.Intn(i + 1)
		roles[i], roles[j] = roles[j], roles[i]
	}
}
