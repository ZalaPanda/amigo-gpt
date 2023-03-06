<script lang="ts">
    import { Configuration, OpenAIApi, type ChatCompletionRequestMessage, type CreateCompletionResponseUsage } from "openai";
    import { marked } from "marked";

    let system = "You are a sarcastic assistant. You use markdown. Your favorite languages are JS, GO and C#.";
    let messages = [] as ChatCompletionRequestMessage[];
    // let messages = [
    //     { role: "user", content: 'Write a tiny program the outputs the "test" word.' },
    //     {
    //         role: "assistant",
    //         content:
    //             'Sure :)\n```javascript\nconsole.log("test");\n```\n```go\npackage main\n\nimport "fmt"\n\nfunc main() {\n    fmt.Println("test")\n}\n```\n```csharp\nusing System;\n\npublic class Program\n{\n    public static void Main()\n    {\n        Console.WriteLine("test");\n    }\n}\n``` \n\nAll three languages will output the "test" word when executed.',
    //     },
    // ] as ChatCompletionRequestMessage[];
    let usage = {
        completion_tokens: 98,
        prompt_tokens: 146,
        total_tokens: 244,
    } as CreateCompletionResponseUsage;

    const configuration = new Configuration({ apiKey: localStorage.getItem("OPENAI_API_KEY") });
    const openai = new OpenAIApi(configuration);
    const onInput = (event: KeyboardEvent & { currentTarget: HTMLTextAreaElement }) => {
        if (event.key !== "Enter" || event.shiftKey) return;
        event.preventDefault();
        const { currentTarget: input } = event;
        console.log(input.value);
        if (input.value) load(input.value);
        input.value = "";
        input.select();
    };
    const load = async (prompt: string) => {
        messages = messages.concat({ role: "user", content: prompt });
        const response = await openai.createChatCompletion(
            {
                model: "gpt-3.5-turbo",
                messages: [{ role: "system", content: system }, ...messages.map(({ role, content }) => ({ role, content }))],
            },
            {
                headers: { "User-Agent": null },
            }
        );
        messages = messages.concat(response.data.choices[0].message);
        usage = response.data.usage;
        console.log(response);
    };
</script>

<section>
    {#each messages as message}
        <article class={message.role}>
            {@html marked.parse(message.content, { breaks: true, xhtml: true })}
        </article>
    {/each}
</section>
<textarea on:keydown={onInput}>Write a tiny program the outputs the "test" word.</textarea>
{#if usage}<aside>{usage.total_tokens}</aside>{/if}

<style lang="less">
    aside {
        position: relative;
        right: 0;
        bottom: 0;
        margin: 10px;
        display: inline-flex;
    }
    textarea {
        font-family: inherit;
        font-size: 1em;
        border: none;
        outline: none;
        // height: 30px;
        // line-height: 30px;
        // padding: 0 10px;
        border-radius: 10px;
        padding: 10px 14px;
        background-color: rgba(240, 240, 240, 1);
        -webkit-font-smoothing: antialiased;
        margin: 5px;
    }
    section {
        text-align: left;
        display: flex;
        align-items: flex-start;
        flex-direction: column;
        & > article {
            :global(& > p) {
                margin: 2px 0;
            }
            :global(& > pre) {
                padding: 10px;
                background-color: #333333aa;
                border-radius: 10px;
                // border: 1px solid #33333333;
            }
            &.user {
                background-color: #2a9d8f;
            }
            &.assistant {
                background-color: #1e293b;
            }
            // display: inline-block;
            position: relative;
            // font-weight: bold;
            // color: #ffffff;
            // border: 1px solid #ffffff;
            border-radius: 10px;
            padding: 10px 14px;
            margin: 5px;
        }
    }
</style>
