defmodule Bot.Commands do
  alias Bot.Commands

  @prefix "!"

  # If the message comes from our bot, we don't want a response
  def handle(%{author: %{id: @bot_id}}), do: :noop

  def handle(msg = %{context: @prefix <> content}) do
    content
    |> String.trim()
    |> Sring.split(" ", parts: 3)
    |> execute(msg)
  end

  def handle(_), do: :noop

  defp execute(["hello"], msg) do
    create_message(msg.channel_id, "Hello to you too!")
  end

  defp execute(["challenges", challenge_name], msg) do
    challenge = Command.Challenges.get_challenge(challenge_name)
    create_message(msg.channel_id, challenge)
  end

  defp execute(_, msg) do
    create_message(msg.channel_id, "NO CAN DO")
  end

  defp create_message(channel_id, message) do
    NostrumApi.create_message(channel_id, message)
  end
end
