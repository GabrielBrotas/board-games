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
			Roles:    []string{"Astronauta", "Engenheiro de Voo", "Comandante", "Biólogo", "Físico", "Médico", "Especialista em Comunicações", "Operador de Carga", "Técnico de Manutenção", "Cientista de Dados", "Psicólogo", "Meteorologista", "Engenheiro Robótico", "Químico", "Astrônomo", "Geólogo", "Nutricionista", "Piloto", "Gerente de Projeto", "Visitante"},
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
			Roles:    []string{"Gerente", "Caixa", "Segurança", "Analista Financeiro", "Atendente", "Técnico de TI", "Economista", "Auditor", "Advogado", "Consultor de Investimentos", "Recepcionista", "Estagiário", "Analista de Crédito", "Operador de Câmbio", "Especialista em Riscos", "Administrador de Patrimônio", "Limpeza", "Gerente de Contas", "Corretor de Seguros", "Agente Comercial"},
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
			Roles:    []string{"Croupier", "Segurança", "Gerente", "Bartender", "Caixa", "Dealer de Poker", "Cantor", "Limpeza", "Administrador Financeiro", "Atendente de Estacionamento", "Recepcionista", "Contador", "Supervisor de andar", "Garçom", "Vigilante", "Especialista em TI", "Técnico de Som e Luz", "Músico", "Gerente de Marketing", "Fotógrafo"},
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
