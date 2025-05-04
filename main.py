from langchain_ollama.llms import OllamaLLM
from langchain_core.prompts import ChatPromptTemplate


def main():
    model = OllamaLLM(model="llama3.2")

    template = """
    You are a dungeon master for a pen and paper session in a fantasy environment
    The description of the character: {description}
    Provide the information as xml file with the following information
    <agility>
    <strength>
    <stamina>
    <intelligence>
    <charisma>

    The question: {question}
    """
    prompt = ChatPromptTemplate.from_template(template)
    chain = prompt | model

    description = """
    Balgrog is a dwarf engineer who has his crossbow strongarm with him. He always hits his target from a long distance
    and because his crossbow is extremely big and strong he is able to strike from a wide distance.
    He loves to tinker in his workshop and creates all kind of different gadgets that help him to fight.
    Because he tinkered so long in his own workshop he is being used to be alone, so in group scenarios
    his actions fail sometimes beacuse he struggles with people.
    """

    question = """
    Provide status values an gear the character has.
    """

    result = chain.invoke({"description": [description], "question": question})
    print(result)


if __name__ == '__main__':
    main()
