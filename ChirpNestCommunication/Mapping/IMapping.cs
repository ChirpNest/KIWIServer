
namespace ChirpNestCommunication.Mapping
{
    public interface IMapping<in TFrom, TTo>
        where TFrom : class
        where TTo : class
    {
        TTo Map(TFrom source);
    }
}