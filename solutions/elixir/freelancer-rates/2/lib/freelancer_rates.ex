defmodule FreelancerRates do
  @spec daily_rate(hourly_rate :: integer) :: float
  def daily_rate(hourly_rate) do
    hourly_rate * hours_per_day()
  end

  @spec apply_discount(before_discount :: integer, discount :: number) :: float
  def apply_discount(before_discount, discount) do
    # before_discount - (before_discount * percent_to_float(discount))
    percent_to_float(discount)
    |> Kernel.*(before_discount)
    |> Kernel.*(-1)
    |> Kernel.-(before_discount)
  end

  @spec monthly_rate(hourly_rate :: integer, discount :: number) :: integer
  def monthly_rate(hourly_rate, discount) do
    Kernel.ceil(apply_discount(daily_rate(hourly_rate), discount) * days_per_month())
  end

  @spec days_in_budget(budget :: integer, hourly_rate :: integer, discount :: number) :: float
  def days_in_budget(budget, hourly_rate, discount) do
    Float.floor(budget / apply_discount(daily_rate(hourly_rate), discount), 1)
  end

  @spec hours_per_day() :: float
  defp hours_per_day, do: 8.0

  @spec days_per_month() :: integer
  defp days_per_month, do: 22

  @spec percent_to_float(percent :: number) :: float
  defp percent_to_float(percent), do: percent / 100
end
