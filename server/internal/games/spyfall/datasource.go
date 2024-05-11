package spyfall

import "math/rand"

type iLocationAndRoles struct {
	Location string
	slug     string
	Roles    []string
}

func getLocationAndRoles() []iLocationAndRoles {
	return []iLocationAndRoles{
		{
			Location: "Hospital",
			Roles:    []string{"Médico", "Enfermeiro", "Cirurgião", "Segurança", "Recepcionista", "Administrador", "Técnico de Laboratório", "Radiologista", "Farmacêutico", "Psicólogo", "Assistente Social", "Nutricionista", "Fisioterapeuta", "Anestesista", "Pediatra", "Dermatologista", "Cardiologista", "Ortopedista", "Visitante", "Paciente"},
			slug:     "hospital",
		},
		{
			Location: "Estação Espacial",
			Roles:    []string{"Astronauta", "Engenheiro de Voo", "Comandante", "Biólogo", "Físico", "Médico", "Especialista em Comunicações", "Operador de Carga", "Técnico de Manutenção", "Cientista de Dados", "Psicólogo", "Meteorologista", "Engenheiro Robótico", "Químico", "Astrônomo", "Turista", "Nutricionista", "Piloto", "Gerente de Projeto", "Visitante"},
			slug:     "estacao-espacial",
		},
		{
			Location: "Supermercado",
			Roles:    []string{"Caixa", "Gerente", "Repositor", "Segurança", "Açougueiro", "Padeiro", "Atendente de Frios", "Estoquista", "Empacotador", "Fiscal de Loja", "Atendente de Sac", "Limpeza", "Administrador", "Nutricionista", "Técnico de Informática", "Analista de Estoque", "Promotor de Vendas", "Balconista", "Confeiteiro", "Jovem Aprendiz"},
			slug:     "supermercado",
		},
		{
			Location: "Submarino",
			Roles:    []string{"Capitão", "Primeiro Oficial", "Engenheiro Chefe", "Operador de Sonar", "Mergulhador", "Médico", "Técnico em Eletrônica", "Operador de Comunicações", "Cientista Marinho", "Técnico em Mecânica", "Eletricista", "Analista de Sistemas", "Operador de Torpedo", "Especialista em Navegação", "Meteorologista", "Operador de Radar", "Físico Nuclear", "Visitante", "Diretor de Engenharia Submarina", "Gerente de Projetos"},
			slug:     "submarino",
		},
		{
			Location: "Banco",
			Roles:    []string{"Gerente", "Caixa", "Segurança", "Analista Financeiro", "Atendente", "Técnico de TI", "Economista", "Ladrão", "Advogado", "Consultor de Investimentos", "Recepcionista", "Estagiário", "Analista de Crédito", "Operador de Câmbio", "Especialista em Riscos", "Administrador de Patrimônio", "Limpeza", "Gerente de Contas", "Corretor de Seguros", "Agente Comercial"},
			slug:     "banco",
		},
		{
			Location: "Escola",
			Roles:    []string{"Professor", "Diretor", "Coordenador Pedagógico", "Secretário", "Psicólogo Escolar", "Nutricionista", "Técnico de Informática", "Porteiro", "Limpeza", "Monitor de Laboratório", "Monitor de Educação Física", "Cozinheira", "Conselheiro Tutelar", "Pais de Aluno", "Vigia", "Motorista de Ônibus Escolar", "Marketing", "Administrador", "Professor de Ingles", "Professor de Matematica"},
			slug:     "escola",
		},
		{
			Location: "Circo",
			Roles:    []string{"Palhaço", "Malabarista", "Domador de Leões", "Acrobata", "Mágico", "Vendedor de Ingressos", "Trapezista", "Equilibrista", "Músico", "Iluminador", "Costureira de Figurinos", "Segurança", "Administrador", "Bilheteiro", "Fotógrafo", "Diretor Artístico", "Coreógrafo", "Maquiador", "Veterinário", "Limpeza"},
			slug:     "circo",
		},
		{
			Location: "Restaurante",
			Roles:    []string{"Garçom", "Cozinheiro", "Barman", "Recepcionista", "Gerente", "Lavador de Pratos", "Churrasqueiro", "Pizzaiolo", "Confeiteiro", "Segurança", "Limpeza", "Caixa", "Decorador", "Estoquista", "Nutricionista", "Chef de Cozinha", "Entregador", "Auxiliar de Cozinha", "Auxiliar de Limpeza", "Marketing"},
			slug:     "restaurante",
		},
		{
			Location: "Teatro",
			Roles:    []string{"Ator", "Diretor", "Cenógrafo", "Iluminador", "Sonoplasta", "Maquiador", "Figurinista", "Bilheteiro", "Coreógrafo", "Técnico de Som", "Técnico de Luz", "Produtor", "Assistente de Palco", "Recepcionista", "Limpeza", "Fotógrafo", "Segurança", "Administrador", "Pintor de Cenários", "Cameraman"},
			slug:     "teatro",
		},
		{
			Location: "Aeroporto",
			Roles:    []string{"Piloto", "Comissário de Bordo", "Agente de Check-in", "Segurança", "Controlador de Tráfego Aéreo", "Mecânico de Aviação", "Carregador de Bagagem", "Limpeza", "Gerente de Operações", "Técnico em Meteorologia", "Operador de Raio X", "Chefe de Cabine", "Recepcionista de Sala VIP", "Manobrista", "Guia de Turismo", "Motorista de Ônibus Aeroportuário", "Agente de Imigração", "Administrador de Aeroporto", "Técnico de Informática", "Profissional de Logistica"},
			slug:     "aeroporto",
		},
		{
			Location: "Zoológico",
			Roles:    []string{"Biólogo", "Veterinário", "Guia Turístico", "Cuidador de Animais", "Segurança", "Fotógrafo", "Administrador", "Atendente de Loja de Souvenirs", "Jardineiro", "Educador Ambiental", "Operador de Caixa", "Nutricionista de Animais", "Diretor", "Serviços Gerais", "Pesquisador", "Transportador de Animais", "Gestor Ambiental", "Limpeza", "Organizador de Eventos", "Recepcionista"},
			slug:     "zoo",
		},
		{
			Location: "Cassino",
			Roles:    []string{"Croupier", "Segurança", "Gerente", "Bartender", "Caixa", "Dealer de Poker", "Apostador", "Limpeza", "Administrador Financeiro", "Atendente de Estacionamento", "Recepcionista", "Contador", "Client", "Garçom", "Vigilante", "Especialista em TI", "Apostador", "Músico", "Gerente de Marketing", "Fotógrafo"},
			slug:     "cassino",
		},
		{
			Location: "Navio Cruzeiro",
			Roles:    []string{"Capitão", "Marinheiro", "Cozinheiro", "Animador de Bordo", "Médico", "Barman", "Recepcionista", "Mecânico", "Segurança", "Garçom", "Instrutor de Esportes Aquáticos", "Fotógrafo", "Guia de Excursões", "Dj", "Operador de Cabine", "Gerente de Eventos", "Camareiro", "Técnico de Som e Luz", "Engenheiro Naval", "Salva-vidas"},
			slug:     "navio-cruzeiro",
		},
		{
			Location: "Parque de Diversões",
			Roles:    []string{"Operador de Brinquedo", "Segurança", "Bilheteiro", "Cozinheiro", "Vendedor de Alimentos", "Palhaço", "Pintor de Rosto", "Atendente de Loja de Brinquedos", "Gerente de Parque", "Operador de Montanha-Russa", "Escultor de Balões", "Artista de Circo", "Técnico de Efeitos Especiais", "Especialista em Segurança de Brinquedos", "Artista de Maquiagem Teatral", "Diretor Artistico", "Designer de Parques Temáticos", "Operador de Carrosel", "Coordenador de Higiene e Segurança", "Consultor de Experiência do Visitante"},
			slug:     "parque-diversoes",
		},
		{
			Location: "Museu",
			Roles:    []string{"Guia Turístico", "Segurança", "Designer de Exposições", "Pesquisador", "Conservador de Arte", "Bibliotecário", "Restaurador de Documentos", "Arqueólogo", "Especialista em Digitalização", "Especialista em Conservação de Materiais", "Especialista em Restauração de Esculturas", "Educador de Museu", "Diretor de Museu", "Guia de Museu", "Recepcionista", "Coordenador de Eventos", "Gerente de Loja de Presentes", "Administrador de Coleções", "Coordenador de Pesquisa", "Fotógrafo de Arte"},
			slug:     "museu",
		},
		{
			Location: "Estúdio de TV",
			Roles:    []string{"Apresentador", "Cinegrafista", "Figurinista", "Maquiador", "Engenheiro de Som", "Editor de Vídeo", "Iluminador", "Ator", "Roteirista", "Assistente de Produção", "Técnico de Informática", "Operador de Câmera", "Jornalista", "Diretor de TV", "Assistente de Direção", "Operador de Teleprompter", "Coordenador de Elenco", "Produtor", "Técnico de Transmissão", "Gerente de Programação"},
			slug:     "estudio-tv",
		},
		// ---
		{
			Location: "Avião",
			Roles:    []string{"Piloto", "Copiloto", "Comissário de Bordo", "Mecânico de Aviação", "Engenheiro Aeronáutico", "Terrorista", "Coordenador de Voo", "Agente de Bagagem", "Técnico de Segurança de Voo", "Instrutor de Voo", "Controlador de Tráfego Aéreo", "Atendente de Serviços ao Passageiro", "Chefe de Cabine", "Agente de Limpeza de Aeronaves", "Passageiro Primeira Classe", "Analista de Operações Aéreas", "Técnico em Aviônica", "Especialista em Navegação Aérea", "Passageiro Classe Economica", "Operador de Rampa"},
			slug:     "aviao",
		},
		{
			Location: "Praia",
			Roles:    []string{"Salva-vidas", "Instrutor de Surf", "Vendedor Ambulante", "Gerente de Quiosque", "Monitor Ambiental", "Fotógrafo de Praia", "Instrutor de Mergulho", "Operador de Passeio de Barco", "Banhista", "Instrutor de Vôlei de Praia", "Operador de Aluguel de Equipamentos", "Artista de Areia", "Guarda Costeira", "Jogador de Altinha", "Surfista", "Massagista", "Ladrão", "Guia Turístico Local", "Promotor de Shows na Praia", "Encarregado de Limpeza de Praia"},
			slug:     "praia",
		},
		{
			Location: "Cinema",
			Roles:    []string{"Bilheteiro", "Operador de Projeção", "Cliente", "Gerente de Cinema", "Técnico de Som", "Agente de Limpeza", "Segurança", "Programador de Filmes", "Assistente de Marketing", "Host de Estreias", "Designer Gráfico", "Analista de Atendimento ao Cliente", "Operador de Cabine", "Especialista em Tecnologia de Projeção", "Vendedor", "Operador de Máquina de Pipoca", "Coordenador de Eventos", "Assistente de Direção", "Operador de Máquina de Refrigerante", "Especialista em Efeitos Visuais"},
			slug:     "cinema",
		},
		{
			Location: "Base Militar",
			Roles:    []string{"Comandante da Base", "Oficial de Operações", "Médico Militar", "Engenheiro Militar", "Sargento de Armas", "Especialista em Logística", "Controlador de Tráfego Aéreo Militar", "Técnico em Comunicações", "Instrutor de Combate", "Mecânico de Veículos Militares", "Operador de Radar", "Soldado", "Cozinheiro Militar", "Operador de Drones", "Especialista em Armamento", "Psicólogo Militar", "Sargento", "Recruta", "Coronel", "Sniper"},
			slug:     "base-militar",
		},
		{
			Location: "Spa",
			Roles:    []string{"Massoterapeuta", "Esteticista", "Recepcionista", "Gerente de Spa", "Instrutor de Yoga", "Terapeuta Holístico", "Nutricionista", "Acupunturista", "Terapeuta de Reflexologia", "Especialista em Aromaterapia", "Coordenador de Atendimento ao Cliente", "Terapeuta de Reiki", "Instrutor de Pilates", "Gerente de Produtos de Spa", "Especialista em Terapias com Água", "Coordenador de Eventos de Bem-Estar", "Jardineiro", "Designer de Interiores de Spa", "Especialista em Banhos Terapêuticos", "Terapeuta Capilar"},
			slug:     "spa",
		},
		{
			Location: "Trem",
			Roles:    []string{"Maquinista", "Condutor", "Fiscal de Trem", "Atendente de Bordo", "Técnico de Manutenção de Trilhos", "Operador de Estação", "Agente de Bilheteria", "Engenheiro Ferroviário", "Chefe de Serviço de Passageiros", "Operador de Sala de Controle", "Segurança Ferroviária", "Coordenador de Logística", "Agente de Limpeza", "Atendente de Bar no Trem", "Técnico em Sistemas de Sinalização", "Operador de Carga", "Assistente de Atendimento ao Cliente", "Gerente de Operações Ferroviárias", "Engenheiro de Sistemas de Trem", "Especialista em Segurança Ferroviária"},
			slug:     "trem",
		},
		{
			Location: "Delegacia",
			Roles:    []string{"Delegado", "Investigador", "Agente de Polícia", "Escrivão", "Coordenador de Operações", "Técnico em Informática Forense", "Analista Criminal", "Operador de Câmeras de Segurança", "Perito Forense", "Administrador de Sistema Penitenciário", "Psicólogo Policial", "Especialista em Relações Comunitárias", "Instrutor de Armas de Fogo", "Operador de Comunicações", "Motorista Policial", "Assistente Administrativo", "Guarda de Cela", "Negociador de Reféns", "Agente de Trânsito", "Coordenador de Programas de Prevenção ao Crime"},
			slug:     "delegacia",
		},
		{
			Location: "Oficina",
			Roles:    []string{"Mecânico de Automóveis", "Eletricista Automotivo", "Pintor de Carros", "Técnico de Ar Condicionado Automotivo", "Lanterneiro", "Gerente de Oficina", "Recepcionista de Oficina", "Estoquista de Peças", "Técnico de Diagnóstico de Veículos", "Vendedor de Peças Automotivas", "Lavador de Carros", "Técnico em Alinhamento e Balanceamento", "Especialista em Restauração de Veículos Antigos", "Técnico em Eletrônica Automotiva", "Coordenador de Serviços", "Assistente de Garantia", "Inspetor de Qualidade", "Operador de Máquinas Ferramenta", "Consultor Técnico Automotivo", "Especialista em Pneus"},
			slug:     "oficina",
		},
		{
			Location: "Estádio de Futebol",
			Roles:    []string{"Gerente de Estádio", "Segurança de Estádio", "Coordenador de Eventos Esportivos", "Relações Públicas do Clube", "Técnico de Futebol", "Jogador de Futebol", "Fisioterapeuta Esportivo", "Narrador Esportivo", "Analista de Desempenho", "Médico Esportivo", "Operador de Câmera", "Diretor de Marketing Esportivo", "Vendedor de Ingressos", "Coordenador de Hospitalidade", "Administrador de Concessões Alimentícias", "Diretor de Operações de Jogo", "Engenheiro de Manutenção de Estádio", "Supervisor de Limpeza", "Organizador de Torcida", "Agente de Atendimento ao Cliente"},
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

func shuffleRoles(roles []string) {
	for i := range roles {
		j := rand.Intn(i + 1)
		roles[i], roles[j] = roles[j], roles[i]
	}
}
