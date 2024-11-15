from abc import ABC, abstractmethod
from typing import List, Optional

from connection_interface import ConnectionInterface


class PacketInterface(ABC):
    @abstractmethod
    def sender(self) -> ConnectionInterface:
        """Returns the object representing the sender of the packet."""
        pass

    @abstractmethod
    def payload(self) -> List[bytearray]:
        """Returns the payload of the packet as a byte array."""
        pass

    @abstractmethod
    def set_payload(self, payload: List[bytearray]) -> None:
        """Sets the payload of the packet."""
        pass

    @abstractmethod
    def rmc_message(self) -> Optional[RMCMessage]:
        """Returns the RMCMessage associated with the packet, if any."""
        pass

    @abstractmethod
    def set_rmc_message(self, message: RMCMessage) -> None:
        """Assigns an RMCMessage object to the packet."""
        pass