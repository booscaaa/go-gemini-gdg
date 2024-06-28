# Sistema de Gerenciamento de Produto com Integração Alexa

Utiliza a Alexa para interagir com usuários e integrar com diversos serviços e bases de dados. Este sistema é projetado para fornecer uma maneira eficiente e interativa de acessar e atualizar informações de produtos através de comandos de voz.

## Visão Geral do Projeto

Este projeto incorpora uma habilidade customizada da Alexa chamada "Prato Rápido", que permite aos usuários consultar informações de produtos rapidamente ou executar atualizações nos dados de produtos armazenados em nosso banco de dados central Postgres. A habilidade interage com vários serviços externos e internos para fornecer uma experiência de usuário rica e responsiva.

## Características Principais

- **Integração com Alexa**: Utilize comandos de voz para interagir diretamente com o sistema sem a necessidade de interfaces tradicionais.
- **Atualização Dinâmica de Dados**: O sistema se conecta a múltiplos serviços externos, como Google e Gemini, para enriquecer e atualizar os dados dos produtos.
- **Automatização de Tarefas**: Atualizações automáticas do catálogo de produtos nos sites Didio's e Potatos a cada 30 minutos, garantindo que a informação esteja sempre atualizada e precisa.

## Objetivo do Sistema

O principal objetivo deste sistema é facilitar o acesso e a gestão de informações de produtos de forma rápida e eficiente, utilizando a tecnologia de voz da Alexa para melhorar a experiência do usuário e aumentar a eficácia operacional.

## Como Contribuir

Estamos abertos a contribuições! Se você tem interesse em melhorar o sistema ou sugerir novas funcionalidades, sinta-se à vontade para criar um pull request ou abrir um issue.

## Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo LICENSE para mais detalhes.


# Fluxo de Interação com Alexa

![Arquitetura do Software](/assets/diagrama_de_sequencia.svg)

Este diagrama de sequência demonstra como um usuário interage com a Alexa e outros componentes de backend para realizar tarefas específicas por meio de uma habilidade da Alexa.

## Fluxo de Interação

1. **Ator**: O usuário inicia a interação ativando a Alexa.
2. **Alexa**: Recebe o comando de voz do usuário e ativa a habilidade desejada.
3. **Skill**: A habilidade da Alexa, após ser ativada, consulta um endpoint específico.
4. **Webhook**: Este endpoint realiza a busca de produtos no banco de dados.
5. **Postgres**: O banco de dados retorna uma lista de produtos para o webhook.
6. **Webhook**: O webhook envia a lista de produtos de volta para a habilidade da Alexa, que então pede uma análise mais detalhada.
7. **LLM (Language Learning Model)**: A análise é feita pelo modelo de aprendizado de línguas, que processa e retorna o texto analisado.
8. **Skill**: A habilidade da Alexa, agora com o texto processado, envia a resposta final para a Alexa.
9. **Alexa**: Apresenta a resposta ao usuário.

## Detalhes do Fluxo

- **Início da Interação**: O usuário ativa a Alexa e faz um comando de voz.
- **Consulta de Produtos**: A habilidade consulta um endpoint que busca produtos no banco de dados.
- **Análise de Produtos**: A lista de produtos é analisada por um modelo de aprendizado de línguas, que gera um texto com base nos dados dos produtos.
- **Resposta ao Usuário**: Alexa fornece a resposta processada ao usuário.

## Conclusão

Este fluxo demonstra a integração entre Alexa e sistemas de backend para processar informações e responder ao usuário de maneira eficiente e automatizada.

# Fluxo de Interação do Sistema com Alexa

![Arquitetura do Software](/assets/diagrama_cenarios.svg)

Este diagrama ilustra como um usuário interage com a Alexa para acessar e manipular informações de produtos em diversos serviços e bases de dados.

## Descrição do Fluxo

1. **Ator (Usuário)**: O usuário inicia a interação com um comando de voz, ativando a Alexa.
2. **Alexa, Prato Rápido**: Alexa processa o comando de voz e ativa a habilidade "Prato Rápido".
3. **Skill Prato Rápido**: A habilidade interage com o sistema para buscar ou modificar informações.
4. **Service 2 - Webhook**: O webhook é ativado pela skill e consulta ou envia informações para diversos serviços externos.
5. **Google, Gemini**: Exemplo de serviços externos que podem ser consultados ou notificados pelo webhook.
6. **Postgres**: Banco de dados principal onde as informações dos produtos são armazenadas.
7. **Service1**: Serviço responsável pela atualização regular dos produtos, que acessa o banco de dados Postgres.
8. **Didio's e Potatos Menus**: Sites específicos onde o catálogo de produtos é atualizado ou consultado pelo Service1.

## Funcionalidades Adicionais

- **Resposta da Alexa**: Após o processamento das solicitações, Alexa fornece uma resposta ao usuário baseada nas informações acessadas ou nas ações realizadas.
- **Ciclo de Atualização**: Service1 realiza atualizações periódicas no banco de dados a cada 30 minutos, garantindo que os dados estejam sempre atualizados.

## Conclusão

Este fluxo mostra a integração entre uma interface de voz com Alexa e múltiplos back-ends e serviços externos para uma gestão dinâmica e atualizada de informações de produtos.

# Arquitetura do Software

![Arquitetura do Software](/assets/diagrama_software.svg)

Este diagrama representa a arquitetura da nossa aplicação de software, que é focada na interação entre diversos componentes relacionados ao gerenciamento de produtos.

## Componentes

### Entradas do Usuário

- **ProductController**: Lida com solicitações HTTP para interfaces web.
- **ProductCLI**: Gerencia entradas via linha de comando.

Ambos os componentes enviam comandos para o `ProductUseCase`.

### Lógica Central

- **ProductUseCase**: Atua como o manipulador central da lógica de negócios para operações relacionadas a produtos. Interage com as interfaces de entrada e os repositórios de dados para executar operações de busca e armazenamento de dados.

### Repositórios de Dados

- **ProductDatabaseRepository**: Gerencia consultas a bancos de dados estruturados.
- **ProductScraperRepository**: Envolve-se com a extração de dados de fontes externas (web scraping).
- **ProductLLMRepository**: Lida com interações com modelos de linguagem ou aprendizado de máquina, como para geração de descrições de produtos.

### Casos de Uso Específicos

- **ProductScraperUseCase**: Gerencia a lógica e o fluxo de coleta de dados de fontes externas. Interage diretamente com `ProductScraperRepository` e `ProductLLMRepository`.

## Fluxo de Informações

- `ProductUseCase` tem acesso direto a todos os três repositórios, indicando uma centralização da lógica de manipulação de dados.
- `ProductScraperUseCase` interage especificamente com `ProductScraperRepository` e `ProductLLMRepository`, o que sugere que os dados coletados são processados ou aprimorados antes de seu uso ou armazenamento.

## Conclusão

Esta arquitetura promove a separação de preocupações, dividindo claramente as responsabilidades entre os componentes e melhorando a manutenção e escalabilidade da aplicação.




