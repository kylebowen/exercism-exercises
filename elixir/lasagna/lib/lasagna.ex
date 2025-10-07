defmodule Lasagna do
  @spec expected_minutes_in_oven() :: integer
  def expected_minutes_in_oven(), do: 40

  @spec remaining_minutes_in_oven(elapsed_minutes_in_oven :: integer) :: integer
  def remaining_minutes_in_oven(elapsed_minutes_in_oven) do
    expected_minutes_in_oven() - elapsed_minutes_in_oven
  end

  @spec preparation_time_in_minutes(layer_count :: integer) :: integer
  def preparation_time_in_minutes(layer_count) do
    minutes_per_layer = 2
    layer_count * minutes_per_layer
  end

  @spec total_time_in_minutes(layer_count :: integer, elapsed_minutes :: integer) :: integer
  def total_time_in_minutes(layer_count, elapsed_minutes) do
    preparation_time_in_minutes(layer_count) + elapsed_minutes
  end

  @spec alarm() :: String.t()
  def alarm(), do: "Ding!"
end
