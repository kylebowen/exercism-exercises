defmodule LanguageList do
  @type language :: String.t()
  @type language_list :: list(language()) | []

  @spec new() :: language_list()
  def new(), do: []

  @spec add(language_list(), language()) :: language_list()
  def add(list, language), do: [language | list]

  @spec remove(language_list()) :: language_list()
  def remove(list), do: tl(list)

  @spec first(language_list()) :: language()
  def first(list), do: hd(list)

  @spec count(language_list()) :: non_neg_integer()
  def count(list), do: length(list)

  @spec functional_list?(language_list()) :: boolean()
  def functional_list?(list), do: "Elixir" in list
end
